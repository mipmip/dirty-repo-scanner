package ui

import (
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mipmip/dirty-repo-scanner/scanner"
)

const (
	viewRepo   = 0
	viewStatus = 1
	viewLog    = 2
)

// Message types

type scanMsg struct {
	repositories scanner.MultiGitStatus
	err          error
}

type logMsg string

// logWriter sends log output as tea messages to the program.
type logWriter struct {
	program *tea.Program
}

func (w logWriter) Write(p []byte) (n int, err error) {
	w.program.Send(logMsg(string(p)))
	return len(p), nil
}

// Styles

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Padding(0, 1)

	selectedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("2")). // green
			Foreground(lipgloss.Color("0")). // black
			Width(0)                         // set dynamically

	normalStyle = lipgloss.NewStyle()

	activeBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("2")) // green

	inactiveBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("240")) // gray

	modalStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("1")). // red
			Padding(1, 2).
			Align(lipgloss.Center)
)

type model struct {
	config            *scanner.Config
	ignoreDirErrors   bool
	repositories      scanner.MultiGitStatus
	repoPaths         []string
	cursor            int
	activeView        int
	scanning          bool
	err               error
	spinner           spinner.Model
	statusViewport    viewport.Model
	logViewport       viewport.Model
	logContent        string
	width             int
	height            int
	program           *tea.Program
}

func newModel(config *scanner.Config, ignoreDirErrors bool) model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return model{
		config:          config,
		ignoreDirErrors: ignoreDirErrors,
		scanning:        true,
		spinner:         s,
		statusViewport:  viewport.New(0, 0),
		logViewport:     viewport.New(0, 0),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.doScan(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.recalcLayout()

	case tea.KeyMsg:
		if m.scanning {
			// Only allow quit during scanning
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			}
			return m, nil
		}

		if m.err != nil {
			// Any key dismisses error
			m.err = nil
			return m, nil
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "s":
			m.scanning = true
			cmds = append(cmds, m.doScan())
		case "e":
			cmd := m.doEdit()
			if cmd != nil {
				return m, cmd
			}
		case "tab":
			m.activeView = (m.activeView + 1) % 3
		case "up":
			if m.activeView == viewRepo {
				if m.cursor > 0 {
					m.cursor--
					m.updateStatusContent()
				}
			} else if m.activeView == viewStatus {
				m.statusViewport.LineUp(1)
			} else if m.activeView == viewLog {
				m.logViewport.LineUp(1)
			}
		case "down":
			if m.activeView == viewRepo {
				if m.cursor < len(m.repoPaths)-1 {
					m.cursor++
					m.updateStatusContent()
				}
			} else if m.activeView == viewStatus {
				m.statusViewport.LineDown(1)
			} else if m.activeView == viewLog {
				m.logViewport.LineDown(1)
			}
		}

	case scanMsg:
		m.scanning = false
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.repositories = msg.repositories
			m.repoPaths = make([]string, 0, len(m.repositories))
			for r := range m.repositories {
				m.repoPaths = append(m.repoPaths, r)
			}
			sort.Strings(m.repoPaths)
			if m.cursor >= len(m.repoPaths) {
				m.cursor = max(0, len(m.repoPaths)-1)
			}
			m.recalcLayout()
			m.updateStatusContent()
		}

	case logMsg:
		m.logContent += string(msg)
		m.logViewport.SetContent(m.logContent)
		m.logViewport.GotoBottom()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *model) recalcLayout() {
	if m.width == 0 || m.height == 0 {
		return
	}
	// Account for borders (2 lines per panel)
	innerWidth := m.width - 2

	statusHeight := m.statusPanelHeight()
	if statusHeight > 0 {
		m.statusViewport.Width = innerWidth
		m.statusViewport.Height = statusHeight
	}

	logHeight := m.logPanelHeight()
	if logHeight > 0 {
		m.logViewport.Width = innerWidth
		m.logViewport.Height = logHeight
	}
}

func (m model) repoPanelHeight() int {
	n := len(m.repoPaths)
	if n == 0 {
		n = 1
	}
	maxH := (m.height - 6) / 2 // leave room for other panels + borders
	if n > maxH {
		return maxH
	}
	return n
}

func (m model) statusPanelHeight() int {
	repoH := m.repoPanelHeight() + 2 // +border
	logH := m.logPanelHeight() + 2    // +border
	remaining := m.height - repoH - logH
	if remaining < 3 {
		return 3
	}
	return remaining - 2 // -border
}

func (m model) logPanelHeight() int {
	return min(10, (m.height-6)/3)
}

func (m *model) updateStatusContent() {
	if len(m.repoPaths) == 0 {
		m.statusViewport.SetContent("")
		return
	}
	currentRepo := m.repoPaths[m.cursor]
	st, ok := m.repositories[currentRepo]
	if !ok || len(st.Status) == 0 {
		m.statusViewport.SetContent("")
		return
	}

	var b strings.Builder
	b.WriteString(" SW\n")
	b.WriteString("-----\n")

	paths := make([]string, 0, len(st.Status))
	for path := range st.Status {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		status := st.Status[path]
		b.WriteString(fmt.Sprintf(" %c%c  %s\n", status.Staging, status.Worktree, path))
	}
	m.statusViewport.SetContent(b.String())
	m.statusViewport.GotoTop()
}

func (m model) doScan() tea.Cmd {
	config := m.config
	ignoreDirErrors := m.ignoreDirErrors
	return func() tea.Msg {
		repos, err := scanner.Scan(config, ignoreDirErrors)
		return scanMsg{repositories: repos, err: err}
	}
}

func (m model) doEdit() tea.Cmd {
	if len(m.repoPaths) == 0 || m.cursor >= len(m.repoPaths) {
		return nil
	}
	currentRepo := m.repoPaths[m.cursor]
	if currentRepo == "" {
		return nil
	}

	cmdStr := strings.Replace(m.config.EditCommand, "%WORKING_DIRECTORY", currentRepo, -1)
	args := strings.Fields(cmdStr)
	if len(args) == 0 {
		return nil
	}

	c := exec.Command(args[0], args[1:]...)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return nil
	})
}

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}
	if m.height < 20 {
		return "Terminal too small. Need at least 20 lines."
	}

	innerWidth := m.width - 2

	// Repo panel
	repoContent := m.renderRepoList(innerWidth)
	repoPanel := m.renderPanel(viewRepo, innerWidth, m.repoPanelHeight(), repoContent)

	// Status panel
	statusPanel := m.renderPanel(viewStatus, innerWidth, m.statusPanelHeight(), m.statusViewport.View())

	// Log panel
	logPanel := m.renderPanel(viewLog, innerWidth, m.logPanelHeight(), m.logViewport.View())

	view := lipgloss.JoinVertical(lipgloss.Left, repoPanel, statusPanel, logPanel)

	// Modal overlays
	if m.scanning {
		modal := modalStyle.Width(20).Render(m.spinner.View() + " Scanning...")
		view = placeOverlay(m.width, m.height, modal, view)
	}
	if m.err != nil {
		errText := fmt.Sprintf("Error: %v", m.err)
		modal := modalStyle.Width(m.width * 3 / 4).Render(errText)
		view = placeOverlay(m.width, m.height, modal, view)
	}

	return view
}

func (m model) renderRepoList(width int) string {
	if len(m.repoPaths) == 0 {
		return "No dirty repositories found."
	}

	var b strings.Builder
	h := m.repoPanelHeight()
	// Calculate scroll offset
	offset := 0
	if m.cursor >= h {
		offset = m.cursor - h + 1
	}

	end := offset + h
	if end > len(m.repoPaths) {
		end = len(m.repoPaths)
	}

	for i := offset; i < end; i++ {
		if i > offset {
			b.WriteString("\n")
		}
		line := m.repoPaths[i]
		if i == m.cursor {
			styled := selectedStyle.Width(width).Render(line)
			b.WriteString(styled)
		} else {
			b.WriteString(normalStyle.Render(line))
		}
	}
	return b.String()
}

func (m model) renderPanel(view int, width int, height int, content string) string {
	var title string
	switch view {
	case viewRepo:
		title = " Repositories "
	case viewStatus:
		title = " Status "
	case viewLog:
		title = " Log "
	}

	borderColor := lipgloss.Color("240")
	if m.activeView == view {
		borderColor = lipgloss.Color("2")
	}

	// Build border with title embedded in top line
	border := lipgloss.RoundedBorder()
	titleStyled := lipgloss.NewStyle().Foreground(borderColor).Bold(true).Render(title)
	topBorder := border.TopLeft +
		strings.Repeat(border.Top, 1) +
		titleStyled +
		strings.Repeat(border.Top, max(0, width-lipgloss.Width(title)-2)) +
		border.TopRight

	// Render content box without top border
	boxStyle := lipgloss.NewStyle().
		Border(border).
		BorderTop(false).
		BorderForeground(borderColor).
		Width(width).
		Height(height)

	return topBorder + "\n" + boxStyle.Render(content)
}

// placeOverlay renders a modal centered over the background.
func placeOverlay(width, height int, modal, background string) string {
	return lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		modal,
		lipgloss.WithWhitespaceBackground(lipgloss.NoColor{}),
	)
}

func Run(config *scanner.Config, ignoreDirErrors bool) error {
	m := newModel(config, ignoreDirErrors)
	p := tea.NewProgram(m, tea.WithAltScreen())

	// Set up log writer to pipe into the TUI
	m.program = p
	log.SetOutput(logWriter{program: p})

	_, err := p.Run()
	return err
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

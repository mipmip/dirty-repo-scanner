package ui

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mipmip/dirty-repo-scanner/src/scanner"
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

type diffMsg struct {
	content string
}

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

	navBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("236")).
			Foreground(lipgloss.Color("252"))

	navBarKeyStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("236")).
			Foreground(lipgloss.Color("2")).
			Bold(true)

	diffAddedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))   // green
	diffDeletedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))   // red
	diffHunkStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))   // cyan
	diffMetaStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))   // yellow
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
	fileCursor        int
	filePaths         []string
	diffViewport      viewport.Model
	logVisible        bool
	logShownOnce      bool
	pendingKey        string
	width             int
	height            int
	program           *tea.Program
	inTmux            bool
	version           string
}

func newModel(config *scanner.Config, ignoreDirErrors bool, version string) model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return model{
		config:          config,
		ignoreDirErrors: ignoreDirErrors,
		scanning:        true,
		inTmux:          os.Getenv("TMUX") != "",
		version:         version,
		spinner:         s,
		statusViewport:  viewport.New(0, 0),
		diffViewport:    viewport.New(0, 0),
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

		// Handle pending 'g' chord
		key := msg.String()
		if m.pendingKey == "g" {
			m.pendingKey = ""
			if key == "g" {
				// gg: jump to top
				switch m.activeView {
				case viewRepo:
					m.cursor = 0
					cmds = append(cmds, m.updateStatusContent())
				case viewStatus:
					m.fileCursor = 0
					cmds = append(cmds, m.fetchDiff())
				case viewLog:
					m.logViewport.GotoTop()
				}
				return m, nil
			}
			// Not 'g', fall through to handle key normally
		}

		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "s":
			m.scanning = true
			cmds = append(cmds, m.doScan())
		case "enter":
			cmd := m.doEdit()
			if cmd != nil {
				return m, cmd
			}
		case "e":
			if m.activeView == viewStatus {
				cmd := m.doEditFile()
				if cmd != nil {
					return m, cmd
				}
			}
		case "l":
			m.logVisible = !m.logVisible
			if m.logVisible {
				m.recalcLayout()
				if !m.logShownOnce {
					m.logShownOnce = true
					m.logViewport.GotoBottom()
				}
			} else {
				if m.activeView == viewLog {
					m.activeView = viewRepo
				}
				m.recalcLayout()
			}
		case "tab":
			if m.logVisible {
				m.activeView = (m.activeView + 1) % 3
			} else {
				// Cycle between repo and status only
				if m.activeView == viewRepo {
					m.activeView = viewStatus
				} else {
					m.activeView = viewRepo
				}
			}
			// Fetch diff when entering status panel, clear when leaving
			if m.activeView == viewStatus && len(m.filePaths) > 0 {
				cmds = append(cmds, m.fetchDiff())
			} else {
				m.diffViewport.SetContent("")
			}
		case "g":
			m.pendingKey = "g"
		case "G":
			switch m.activeView {
			case viewRepo:
				if len(m.repoPaths) > 0 {
					m.cursor = len(m.repoPaths) - 1
					cmds = append(cmds, m.updateStatusContent())
				}
			case viewStatus:
				if len(m.filePaths) > 0 {
					m.fileCursor = len(m.filePaths) - 1
					cmds = append(cmds, m.fetchDiff())
				}
			case viewLog:
				m.logViewport.GotoBottom()
			}
		case "pgdown", "ctrl+f":
			half := m.halfPage()
			switch m.activeView {
			case viewRepo:
				m.cursor = min(m.cursor+half, len(m.repoPaths)-1)
				cmds = append(cmds, m.updateStatusContent())
			case viewStatus:
				if len(m.filePaths) > 0 {
					m.fileCursor = min(m.fileCursor+half, len(m.filePaths)-1)
					cmds = append(cmds, m.fetchDiff())
				}
			case viewLog:
				m.logViewport.LineDown(half)
			}
		case "pgup", "ctrl+b":
			half := m.halfPage()
			switch m.activeView {
			case viewRepo:
				m.cursor = max(m.cursor-half, 0)
				cmds = append(cmds, m.updateStatusContent())
			case viewStatus:
				if len(m.filePaths) > 0 {
					m.fileCursor = max(m.fileCursor-half, 0)
					cmds = append(cmds, m.fetchDiff())
				}
			case viewLog:
				m.logViewport.LineUp(half)
			}
		case "up", "k":
			if m.activeView == viewRepo {
				if m.cursor > 0 {
					m.cursor--
					cmds = append(cmds, m.updateStatusContent())
				}
			} else if m.activeView == viewStatus {
				if len(m.filePaths) > 0 && m.fileCursor > 0 {
					m.fileCursor--
					cmds = append(cmds, m.fetchDiff())
				}
			} else if m.activeView == viewLog {
				m.logViewport.LineUp(1)
			}
		case "down", "j":
			if m.activeView == viewRepo {
				if m.cursor < len(m.repoPaths)-1 {
					m.cursor++
					cmds = append(cmds, m.updateStatusContent())
				}
			} else if m.activeView == viewStatus {
				if len(m.filePaths) > 0 && m.fileCursor < len(m.filePaths)-1 {
					m.fileCursor++
					cmds = append(cmds, m.fetchDiff())
				}
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
			cmds = append(cmds, m.updateStatusContent())
		}

	case diffMsg:
		m.diffViewport.SetContent(colorizeDiff(msg.content))
		m.diffViewport.GotoTop()

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
		// Diff viewport takes 60% of inner width
		m.diffViewport.Width = innerWidth * 6 / 10
		m.diffViewport.Height = statusHeight
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
	logH := m.logPanelHeight()
	if logH > 0 {
		logH += 2 // +border only when visible
	}
	remaining := m.height - repoH - logH - 1 // -1 for nav bar
	if remaining < 3 {
		return 3
	}
	return remaining - 2 // -border
}

func (m model) logPanelHeight() int {
	if !m.logVisible {
		return 0
	}
	return min(10, (m.height-6)/3)
}

func (m model) halfPage() int {
	switch m.activeView {
	case viewStatus:
		return max(1, m.statusViewport.Height/2)
	case viewLog:
		return max(1, m.logViewport.Height/2)
	default:
		return max(1, m.repoPanelHeight()/2)
	}
}

// updateStatusContent rebuilds the file list and returns a Cmd to fetch the diff for the first file.
func (m *model) updateStatusContent() tea.Cmd {
	if len(m.repoPaths) == 0 {
		m.filePaths = nil
		m.fileCursor = 0
		m.diffViewport.SetContent("")
		return nil
	}
	currentRepo := m.repoPaths[m.cursor]
	st, ok := m.repositories[currentRepo]
	if !ok || len(st.Status) == 0 {
		m.filePaths = nil
		m.fileCursor = 0
		m.diffViewport.SetContent("")
		return nil
	}

	m.filePaths = make([]string, 0, len(st.Status))
	for path := range st.Status {
		m.filePaths = append(m.filePaths, path)
	}
	sort.Strings(m.filePaths)
	m.fileCursor = 0
	m.diffViewport.SetContent("")
	return nil
}

func (m model) doScan() tea.Cmd {
	config := m.config
	ignoreDirErrors := m.ignoreDirErrors
	return func() tea.Msg {
		repos, err := scanner.Scan(config, ignoreDirErrors)
		return scanMsg{repositories: repos, err: err}
	}
}

func (m model) fetchDiff() tea.Cmd {
	if len(m.repoPaths) == 0 || len(m.filePaths) == 0 {
		return nil
	}
	repoPath := m.repoPaths[m.cursor]
	filePath := m.filePaths[m.fileCursor]

	// Check if file is untracked
	st := m.repositories[repoPath]
	if fs, ok := st.Status[filePath]; ok && fs.Worktree == '?' {
		return func() tea.Msg {
			return diffMsg{content: "Untracked file"}
		}
	}

	return func() tea.Msg {
		cmd := exec.Command("git", "diff", "--", filePath)
		cmd.Dir = repoPath
		out, err := cmd.Output()
		if err != nil || len(out) == 0 {
			// Try staged diff
			cmd = exec.Command("git", "diff", "--cached", "--", filePath)
			cmd.Dir = repoPath
			out, _ = cmd.Output()
		}
		if len(out) == 0 {
			return diffMsg{content: "No diff available"}
		}
		return diffMsg{content: string(out)}
	}
}

func (m model) canEditSelectedFile() bool {
	if len(m.repoPaths) == 0 || len(m.filePaths) == 0 {
		return false
	}
	repoPath := m.repoPaths[m.cursor]
	filePath := m.filePaths[m.fileCursor]
	st := m.repositories[repoPath]
	fs, ok := st.Status[filePath]
	if !ok {
		return false
	}
	// File deleted from worktree — nothing to edit
	return fs.Worktree != 'D'
}

func (m model) doEditFile() tea.Cmd {
	if !m.canEditSelectedFile() {
		return nil
	}
	repoPath := m.repoPaths[m.cursor]
	filePath := m.filePaths[m.fileCursor]
	fullPath := repoPath + "/" + filePath

	editor := os.Getenv("EDITOR")

	if m.inTmux && editor != "" {
		return func() tea.Msg {
			_ = exec.Command("tmux", "display-popup", "-E", "-w", "80%", "-h", "80%", "--", editor, fullPath).Run()
			return nil
		}
	}

	// Try xdg-open first
	if xdgOpen, err := exec.LookPath("xdg-open"); err == nil {
		c := exec.Command(xdgOpen, fullPath)
		return tea.ExecProcess(c, func(err error) tea.Msg {
			return nil
		})
	}

	// Fall back to $EDITOR
	if editor != "" {
		c := exec.Command(editor, fullPath)
		return tea.ExecProcess(c, func(err error) tea.Msg {
			return nil
		})
	}

	return nil
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
	if cmdStr == "" {
		return nil
	}

	if m.inTmux {
		return func() tea.Msg {
			_ = exec.Command("tmux", "display-popup", "-E", "-w", "80%", "-h", "80%", "-d", currentRepo).Run()
			return nil
		}
	}

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
	statusContent := m.renderStatusContent(innerWidth, m.statusPanelHeight())
	statusPanel := m.renderPanel(viewStatus, innerWidth, m.statusPanelHeight(), statusContent)

	// Nav bar
	navBar := m.renderNavBar()

	var view string
	if m.logVisible {
		logPanel := m.renderPanel(viewLog, innerWidth, m.logPanelHeight(), m.logViewport.View())
		view = lipgloss.JoinVertical(lipgloss.Left, repoPanel, statusPanel, logPanel, navBar)
	} else {
		view = lipgloss.JoinVertical(lipgloss.Left, repoPanel, statusPanel, navBar)
	}

	// Modal overlays
	if m.scanning {
		modal := modalStyle.Width(40).Render(m.spinner.View() + " Scanning for dirty repo's...")
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

func (m model) renderFileList(width int, height int) string {
	if len(m.filePaths) == 0 {
		return "No dirty files."
	}

	currentRepo := m.repoPaths[m.cursor]
	st := m.repositories[currentRepo]

	var b strings.Builder
	// Scroll offset
	offset := 0
	if m.fileCursor >= height {
		offset = m.fileCursor - height + 1
	}
	end := offset + height
	if end > len(m.filePaths) {
		end = len(m.filePaths)
	}

	isActive := m.activeView == viewStatus
	for i := offset; i < end; i++ {
		if i > offset {
			b.WriteString("\n")
		}
		path := m.filePaths[i]
		fs := st.Status[path]
		line := fmt.Sprintf(" %c%c  %s", fs.Staging, fs.Worktree, path)
		if isActive && i == m.fileCursor {
			b.WriteString(selectedStyle.Width(width).Render(line))
		} else {
			b.WriteString(normalStyle.Render(line))
		}
	}
	return b.String()
}

func (m model) renderStatusContent(innerWidth int, height int) string {
	listWidth := innerWidth * 4 / 10

	fileList := m.renderFileList(listWidth, height)
	diffContent := m.diffViewport.View()

	return lipgloss.JoinHorizontal(lipgloss.Top, fileList, diffContent)
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

func (m model) renderNavBar() string {
	keys := []struct{ key, action string }{
		{"q", "quit"},
		{"s", "scan"},
		{"enter", "open"},
		{"e", "edit file"},
		{"tab", "switch"},
		{"jk/↑↓", "navigate"},
		{"l", "log"},
		{"pgup/dn", "scroll"},
		{"gg/G", "jump"},
	}

	var left strings.Builder
	for i, k := range keys {
		if i > 0 {
			left.WriteString(navBarStyle.Render("  "))
		}
		left.WriteString(navBarKeyStyle.Render(k.key))
		left.WriteString(navBarStyle.Render(" " + k.action))
	}

	right := navBarStyle.Render("dirty-repo-scanner " + m.version)

	bar := lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		left.String()+strings.Repeat(" ", max(0, m.width-lipgloss.Width(left.String())-lipgloss.Width(right)))+right,
		lipgloss.WithWhitespaceBackground(lipgloss.Color("236")),
	)

	return bar
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

func colorizeDiff(content string) string {
	lines := strings.Split(content, "\n")
	var b strings.Builder
	for i, line := range lines {
		if i > 0 {
			b.WriteString("\n")
		}
		switch {
		case strings.HasPrefix(line, "@@"):
			b.WriteString(diffHunkStyle.Render(line))
		case strings.HasPrefix(line, "+++"), strings.HasPrefix(line, "---"),
			strings.HasPrefix(line, "diff "), strings.HasPrefix(line, "index "):
			b.WriteString(diffMetaStyle.Render(line))
		case strings.HasPrefix(line, "+"):
			b.WriteString(diffAddedStyle.Render(line))
		case strings.HasPrefix(line, "-"):
			b.WriteString(diffDeletedStyle.Render(line))
		default:
			b.WriteString(line)
		}
	}
	return b.String()
}

func Run(config *scanner.Config, ignoreDirErrors bool, version string) error {
	m := newModel(config, ignoreDirErrors, version)
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

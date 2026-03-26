package ui

import (
	"strings"
	"testing"
)

func TestRepoPanelHeight(t *testing.T) {
	tests := []struct {
		name      string
		height    int
		repoCount int
		check     func(t *testing.T, got int)
	}{
		{
			name:      "few repos returns repo count",
			height:    40,
			repoCount: 3,
			check: func(t *testing.T, got int) {
				if got != 3 {
					t.Errorf("got %d, want 3", got)
				}
			},
		},
		{
			name:      "many repos capped below half height",
			height:    40,
			repoCount: 100,
			check: func(t *testing.T, got int) {
				if got >= 20 {
					t.Errorf("got %d, want < 20 (half of height)", got)
				}
			},
		},
		{
			name:      "zero repos returns 1",
			height:    40,
			repoCount: 0,
			check: func(t *testing.T, got int) {
				if got != 1 {
					t.Errorf("got %d, want 1", got)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model{height: tt.height}
			m.repoPaths = make([]string, tt.repoCount)
			got := m.repoPanelHeight()
			tt.check(t, got)
		})
	}
}

func TestLogPanelHeight(t *testing.T) {
	tests := []struct {
		name   string
		height int
	}{
		{"small terminal", 20},
		{"medium terminal", 40},
		{"large terminal", 80},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model{height: tt.height}
			got := m.logPanelHeight()
			if got < 1 {
				t.Errorf("got %d, want >= 1", got)
			}
			if got > 10 {
				t.Errorf("got %d, want <= 10", got)
			}
		})
	}
}

func TestStatusPanelHeight(t *testing.T) {
	tests := []struct {
		name      string
		height    int
		repoCount int
	}{
		{"standard terminal", 40, 5},
		{"small terminal", 20, 2},
		{"many repos", 40, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model{height: tt.height}
			m.repoPaths = make([]string, tt.repoCount)
			got := m.statusPanelHeight()
			if got < 1 {
				t.Errorf("got %d, want >= 1", got)
			}

			// Verify total doesn't exceed terminal height
			repoH := m.repoPanelHeight() + 2
			logH := m.logPanelHeight() + 2
			statusH := got + 2
			total := repoH + logH + statusH
			if total > tt.height {
				t.Errorf("total panel height %d exceeds terminal height %d", total, tt.height)
			}
		})
	}
}

func TestRenderRepoList(t *testing.T) {
	t.Run("all repos visible", func(t *testing.T) {
		m := model{
			height:    40,
			repoPaths: []string{"/home/user/repo1", "/home/user/repo2", "/home/user/repo3"},
			cursor:    0,
		}
		got := m.renderRepoList(80)
		for _, repo := range m.repoPaths {
			if !strings.Contains(got, repo) {
				t.Errorf("output missing repo %q", repo)
			}
		}
	})

	t.Run("cursor repo present", func(t *testing.T) {
		m := model{
			height:    40,
			repoPaths: []string{"/home/user/repo1", "/home/user/repo2", "/home/user/repo3"},
			cursor:    1,
		}
		got := m.renderRepoList(80)
		if !strings.Contains(got, "/home/user/repo2") {
			t.Error("output missing cursor repo /home/user/repo2")
		}
	})

	t.Run("empty repo list", func(t *testing.T) {
		m := model{
			height:    40,
			repoPaths: []string{},
		}
		got := m.renderRepoList(80)
		if !strings.Contains(got, "No dirty repositories found") {
			t.Errorf("expected 'No dirty repositories found', got %q", got)
		}
	})
}

package ui

import (
	"errors"
	"reflect"
	"testing"

	"github.com/mipmip/dirty-repo-scanner/src/scanner"
)

func TestRenderCommand(t *testing.T) {
	tests := []struct {
		name string
		tmpl string
		repo string
		want string
	}{
		{"working directory", "code %WORKING_DIRECTORY", "/home/u/foo", "code /home/u/foo"},
		{"repo name", "tmux new -As %REPO_NAME", "/home/u/foo", "tmux new -As foo"},
		{"both, repeated", "%REPO_NAME %WORKING_DIRECTORY %REPO_NAME", "/a/bar", "bar /a/bar bar"},
		{"no placeholders", "echo hi", "/a/bar", "echo hi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderCommand(tt.tmpl, tt.repo); got != tt.want {
				t.Errorf("renderCommand() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSplitArgs(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    []string
		wantErr bool
	}{
		{"plain", "tmux switch-client -t foo", []string{"tmux", "switch-client", "-t", "foo"}, false},
		{"single quotes with space", "tmux new -As 'my repo'", []string{"tmux", "new", "-As", "my repo"}, false},
		{"double quotes with space", `tmux -c "/path/with spaces"`, []string{"tmux", "-c", "/path/with spaces"}, false},
		{"escaped space", `cd /path/with\ space`, []string{"cd", "/path/with space"}, false},
		{"metachars stay literal", `echo a;b$(x)`, []string{"echo", "a;b$(x)"}, false},
		{"extra whitespace", "  a   b\t c ", []string{"a", "b", "c"}, false},
		{"empty", "", nil, false},
		{"unterminated single", "a 'b", nil, true},
		{"unterminated double", `a "b`, nil, true},
		{"trailing backslash", `a \`, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitArgs(tt.in)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("splitArgs(%q) expected error, got nil", tt.in)
				}
				return
			}
			if err != nil {
				t.Fatalf("splitArgs(%q) unexpected error: %v", tt.in, err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitArgs(%q) = %#v, want %#v", tt.in, got, tt.want)
			}
		})
	}
}

func switchModel(t *testing.T, switchCmd string, run func([]string) error) model {
	t.Helper()
	return model{
		config:     &scanner.Config{SwitchCommand: switchCmd},
		repoPaths:  []string{"/home/u/myrepo"},
		cursor:     0,
		multiplex:  true,
		runCommand: run,
	}
}

func TestDoSwitchSuccessQuits(t *testing.T) {
	var got []string
	m := switchModel(t, "tmux new -As %REPO_NAME -c %WORKING_DIRECTORY", func(argv []string) error {
		got = argv
		return nil
	})
	newM, cmd := m.doSwitch()
	if cmd == nil {
		t.Fatal("expected a quit command on success, got nil")
	}
	if newM.err != nil {
		t.Fatalf("expected no error on success, got %v", newM.err)
	}
	want := []string{"tmux", "new", "-As", "myrepo", "-c", "/home/u/myrepo"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("switch argv = %#v, want %#v", got, want)
	}
}

func TestDoSwitchFailureStaysOpen(t *testing.T) {
	m := switchModel(t, "tmux switch-client -t %REPO_NAME", func(argv []string) error {
		return errors.New("boom")
	})
	newM, cmd := m.doSwitch()
	if cmd != nil {
		t.Error("expected no quit command on failure")
	}
	if newM.err == nil {
		t.Error("expected an error to be surfaced on failure")
	}
}

func TestDoSwitchEmptyCommand(t *testing.T) {
	ran := false
	m := switchModel(t, "   ", func(argv []string) error { ran = true; return nil })
	newM, cmd := m.doSwitch()
	if cmd != nil {
		t.Error("expected no quit command when switch_command is empty")
	}
	if newM.err == nil {
		t.Error("expected an error when switch_command is empty")
	}
	if ran {
		t.Error("command should not run when switch_command is empty")
	}
}

func TestEnterDispatchUsesMultiplex(t *testing.T) {
	// In multiplex mode the switch command runs; the editor path is not taken.
	ran := false
	m := switchModel(t, "true %REPO_NAME", func(argv []string) error { ran = true; return nil })
	if _, cmd := m.doSwitch(); cmd == nil {
		t.Fatal("multiplex switch should return a quit command")
	}
	if !ran {
		t.Error("expected switch command to run in multiplex mode")
	}

	// Without the multiplex flag, doEdit governs Enter and never runs switch_command.
	ran = false
	nm := switchModel(t, "true", func(argv []string) error { ran = true; return nil })
	nm.multiplex = false
	nm.config.EditCommand = "" // empty editor command → doEdit is a no-op
	if cmd := nm.doEdit(); cmd != nil {
		t.Error("expected no edit command with empty edit_command")
	}
	if ran {
		t.Error("switch command must not run when multiplex is off")
	}
}

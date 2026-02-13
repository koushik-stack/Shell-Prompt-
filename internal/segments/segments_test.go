package segments

import (
	"strings"
	"testing"

	"github.com/koushik-stack/Shell-Prompt-/internal/config"
)

func TestGitSegment(t *testing.T) {
	gs := &GitSegment{}

	// Test basic rendering (we're in a git repo now)
	output, err := gs.Render(map[string]interface{}{})
	if err != nil {
		t.Errorf("GitSegment.Render() error = %v", err)
	}
	// Should have output since we're in a git repo
	if output == "" {
		t.Error("Expected non-empty output when in git repo")
	}
	if !strings.Contains(output, " on ") {
		t.Errorf("Expected output to contain ' on ', got %q", output)
	}

	// Test with show_status property
	output, err = gs.Render(map[string]interface{}{
		"show_status": true,
	})
	if err != nil {
		t.Errorf("GitSegment.Render() error = %v", err)
	}
	// Should have output since we're in a git repo
	if output == "" {
		t.Error("Expected non-empty output when in git repo with status")
	}
}

func TestTimeSegment(t *testing.T) {
	ts := &TimeSegment{}

	tests := []struct {
		name      string
		props     map[string]interface{}
		checkFunc func(string) bool
	}{
		{
			name:  "default format",
			props: map[string]interface{}{},
			checkFunc: func(output string) bool {
				return len(output) > 0
			},
		},
		{
			name: "custom format",
			props: map[string]interface{}{
				"format": "2006-01-02",
			},
			checkFunc: func(output string) bool {
				return strings.Contains(output, "-") && len(output) == 10
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := ts.Render(tt.props)
			if err != nil {
				t.Errorf("TimeSegment.Render() error = %v", err)
				return
			}
			if !tt.checkFunc(output) {
				t.Errorf("TimeSegment.Render() = %q, failed check", output)
			}
		})
	}
}

func TestLanguageSegment(t *testing.T) {
	ls := &LanguageSegment{}

	output, err := ls.Render(map[string]interface{}{})
	if err != nil {
		t.Errorf("LanguageSegment.Render() error = %v", err)
	}
	// Should be empty in test environment
	if output != "" {
		t.Logf("LanguageSegment.Render() = %q", output)
	}
}

func TestExitCodeSegment(t *testing.T) {
	ecs := &ExitCodeSegment{}

	// Test without EXIT_CODE env var
	output, err := ecs.Render(map[string]interface{}{})
	if err != nil {
		t.Errorf("ExitCodeSegment.Render() error = %v", err)
	}
	if output != "" {
		t.Errorf("Expected empty output without EXIT_CODE, got %q", output)
	}
}

func TestUsernameSegment(t *testing.T) {
	us := &UsernameSegment{}

	output, err := us.Render(map[string]interface{}{})
	if err != nil {
		t.Errorf("UsernameSegment.Render() error = %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty username")
	}

	// Test with show_host
	output, err = us.Render(map[string]interface{}{
		"show_host": true,
	})
	if err != nil {
		t.Errorf("UsernameSegment.Render() error = %v", err)
	}
	if !strings.Contains(output, "@") {
		t.Errorf("Expected username@hostname format, got %q", output)
	}
}

func TestHostnameSegment(t *testing.T) {
	hs := &HostnameSegment{}

	output, err := hs.Render(map[string]interface{}{})
	if err != nil {
		t.Errorf("HostnameSegment.Render() error = %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty hostname")
	}
}

func TestNewSegment(t *testing.T) {
	tests := []struct {
		segType string
		wantNil bool
	}{
		{"directory", false},
		{"git", false},
		{"time", false},
		{"language", false},
		{"exit_code", false},
		{"username", false},
		{"hostname", false},
		{"unknown", true},
	}

	for _, tt := range tests {
		t.Run(tt.segType, func(t *testing.T) {
			got := New(tt.segType)
			if (got == nil) != tt.wantNil {
				t.Errorf("New(%q) = %v, want nil=%v", tt.segType, got, tt.wantNil)
			}
		})
	}
}

func TestRenderSegments(t *testing.T) {
	cfg := &config.Config{
		Segments: []config.SegmentConfig{
			{Type: "directory"},
			{Type: "time"},
			{Type: "unknown"}, // Should be skipped
		},
	}

	outputs, err := RenderSegments(cfg, "bash")
	if err != nil {
		t.Errorf("RenderSegments() error = %v", err)
	}
	if len(outputs) != 2 {
		t.Errorf("Expected 2 segments, got %d", len(outputs))
	}
	for _, output := range outputs {
		if output == "" {
			t.Error("Expected non-empty segment output")
		}
	}
}

package segments

import (
	"github.com/koushik-stack/Shell-Prompt-/internal/config"
)

// Segment is the interface for all prompt segments
type Segment interface {
	Render(props map[string]interface{}) (string, error)
}

// Registry holds all available segments
var Registry = map[string]func() Segment{
	"directory": func() Segment { return &DirectorySegment{} },
	"git":       func() Segment { return &GitSegment{} },
	"time":      func() Segment { return &TimeSegment{} },
	"language":  func() Segment { return &LanguageSegment{} },
	"exit_code": func() Segment { return &ExitCodeSegment{} },
	"username":  func() Segment { return &UsernameSegment{} },
	"hostname":  func() Segment { return &HostnameSegment{} },
}

// New creates a new segment by type
func New(segType string) Segment {
	if fn, ok := Registry[segType]; ok {
		return fn()
	}
	return nil
}

// RenderSegments renders all segments from configuration
func RenderSegments(cfg *config.Config, shellType string) ([]string, error) {
	var results []string

	for _, segCfg := range cfg.Segments {
		segment := New(segCfg.Type)
		if segment == nil {
			continue
		}

		output, err := segment.Render(segCfg.Props)
		if err != nil {
			continue
		}

		if output != "" {
			results = append(results, output)
		}
	}

	return results, nil
}

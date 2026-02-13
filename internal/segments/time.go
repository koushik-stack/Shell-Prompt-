package segments

import (
	"time"
)

type TimeSegment struct{}

func (ts *TimeSegment) Render(props map[string]interface{}) (string, error) {
	format := "15:04:05"
	if f, ok := props["format"].(string); ok {
		format = f
	}

	return time.Now().Format(format), nil
}

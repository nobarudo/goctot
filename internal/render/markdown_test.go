package render_test

import (
	"bytes"
	"goctot/internal/render"
	"strings"
	"testing"
)

func TestMarkdown(t *testing.T) {
	records := [][]string{
		{"A", "B"},
		{"1", "2"},
	}

	var buf bytes.Buffer
	render.Markdown(&buf, records)

	out := buf.String()

	if !strings.Contains(out, "| A | B |") {
		t.Errorf("markdown header missing")
	}
}

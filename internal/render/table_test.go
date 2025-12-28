package render_test

import (
	"bytes"
	"strings"
	"testing"

	"goctot/internal/render"
)

func TestTable(t *testing.T) {
	records := [][]string{
		{"A", "B"},
		{"1", "2"},
	}

	var buf bytes.Buffer
	render.Table(&buf, records)

	out := buf.String()

	if !strings.Contains(out, "A") {
		t.Errorf("output should contain header")
	}
	if !strings.Contains(out, "1") {
		t.Errorf("output should contain row data")
	}
}


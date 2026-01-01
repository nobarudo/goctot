package render_test

import (
	"bytes"
	"goctot/internal/render"
	"strings"
	"testing"
)

func TestMarkdown(t *testing.T) {
	records := [][]string{
		{"First Name", "Last Name", "SSN"},
		{"John", "Barry", "123456"},
		{"Kathy", "Smith", "687987"},
	}

	var buf bytes.Buffer
	render.Markdown(&buf, records)

	out := buf.String()
	want := `| FIRST NAME | LAST NAME |  SSN   |
|------------|-----------|--------|
| John       | Barry     | 123456 |
| Kathy      | Smith     | 687987 |`

	if !strings.Contains(out, want) {
		t.Errorf("markdown display is corrupted")
		t.Error(out)
	}
}

func TestNoHeaderMarkdown(t *testing.T) {
	records := [][]string{
		{"First Name", "Last Name", "SSN"},
		{"John", "Barry", "123456"},
		{"Kathy", "Smith", "687987"},
	}

	var buf bytes.Buffer
	render.NoHeaderMarkdown(&buf, records)

	out := buf.String()
	want := `| First Name | Last Name | SSN    |
| John       | Barry     | 123456 |
| Kathy      | Smith     | 687987 |`

	if !strings.Contains(out, want) {
		t.Errorf("markdown display is corrupted")
		t.Error(out)
	}
}
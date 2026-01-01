package render_test

import (
	"bytes"
	"strings"
	"testing"

	"goctot/internal/render"
)

func TestTable(t *testing.T) {
	records := [][]string{
		{"First Name", "Last Name", "SSN"},
		{"John", "Barry", "123456"},
		{"Kathy", "Smith", "687987"},
	}

	var buf bytes.Buffer
	render.Table(&buf, records)

	out := buf.String()
	want := `|------------|-----------|--------|
| FIRST NAME | LAST NAME |  SSN   |
|------------|-----------|--------|
| John       | Barry     | 123456 |
| Kathy      | Smith     | 687987 |
|------------|-----------|--------|`

	if !strings.Contains(out, want) {
		t.Errorf("table display is corrupted")
		t.Error(out)
	}
}

func TestNoHeaderTable(t *testing.T) {
		records := [][]string{
		{"First Name", "Last Name", "SSN"},
		{"John", "Barry", "123456"},
		{"Kathy", "Smith", "687987"},
	}

	var buf bytes.Buffer
	render.NoHeaderTable(&buf, records)

	out := buf.String()
	want := `|------------|-----------|--------|
| First Name | Last Name | SSN    |
| John       | Barry     | 123456 |
| Kathy      | Smith     | 687987 |
|------------|-----------|--------|`

	if !strings.Contains(out, want) {
		t.Errorf("table display is corrupted")
		t.Error(out)
	}
}
package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	var buf bytes.Buffer

	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	rootCmd.SetArgs([]string{"version"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}

	out := strings.TrimSpace(buf.String())
	want := "goctot version 0.1.0"

	if out != want {
		t.Errorf("unexpected output: %q", out)
	}
}
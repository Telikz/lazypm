package tests

import (
	"bytes"
	"io"
	"os"
	"testing"

	"lazypm/internal/commands"
)

func TestIssueCmd(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd := commands.GetRootCmd()
	issueCmd, _, err := rootCmd.Find([]string{"issue"})
	if err != nil {
		t.Fatal(err)
	}
	issueCmd.Run(issueCmd, []string{})

	if w.Close() != nil {
		return
	}
	os.Stdout = old

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		return
	}

	expected := "Hello from LazyPM CLI!\n"
	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}

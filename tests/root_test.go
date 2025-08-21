package tests

import (
	"testing"

	"lazypm/internal/commands"
)

func TestRootCmd(t *testing.T) {
	rootCmd := commands.GetRootCmd()

	if rootCmd == nil {
		t.Fatal("rootCmd is nil")
	}

	if rootCmd.Use != "lazypm" {
		t.Errorf("rootCmd.Use = %q, want %q", rootCmd.Use, "lazypm")
	}
}

package main

import (
	"os"
	"path/filepath"
	"testing"
)


func Test_isValidDirectory(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		os.RemoveAll(tempDir)
	})

	t.Run("package.json exists", func(t *testing.T) {
		packageJsonDir := filepath.Join(tempDir, "package.json")
		_, err = os.Create(packageJsonDir)
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			os.RemoveAll(packageJsonDir)
		})

		if err := isValidDirectory(tempDir); err != nil {
			t.Error("Error should be <nil> but...", err)
		}
	})

	t.Run("package.json does not exist", func(t *testing.T) {
		if err := isValidDirectory(tempDir); err == nil {
			t.Error("expected error but nothing happened")
		}
	})
}
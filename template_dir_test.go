package main

import (
	"os"
	"path"
	"strings"
	"testing"
)

func TestTemplateDir(t *testing.T) {
	dir := TemplateDir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Fatalf("'.github' directory for templates must be created")
	}
	if !path.IsAbs(dir) {
		t.Fatalf("TemplateDir() must return absolute path")
	}
	if !strings.Contains(dir, ".github") {
		t.Fatalf("Invalid template dir path: %v", dir)
	}
}

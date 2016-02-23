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

func TestSpecifiedByEnvVar(t *testing.T) {
	w, _ := os.Getwd()
	os.Setenv("DOT_GITHUB_HOME", w)
	defer os.Setenv("DOT_GITHUB_HOME", "")

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

func TestInvalidDir(t *testing.T) {
	os.Setenv("DOT_GITHUB_HOME", "/")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Invalid directory must cause panic")
		}
		defer os.Setenv("DOT_GITHUB_HOME", "")
	}()
	TemplateDir()
}

package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestTemplateDir(t *testing.T) {
	dir := TemplateDir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Fatalf("'.github' directory for templates must be created")
	}
	if !filepath.IsAbs(dir) {
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
	if !filepath.IsAbs(dir) {
		t.Fatalf("TemplateDir() must return absolute path")
	}
	if !strings.Contains(dir, ".github") {
		t.Fatalf("Invalid template dir path: %v", dir)
	}
}

func TestInvalidDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("because I don't know how to specify 'invalid' directory path on Windows")
	}
	os.Setenv("DOT_GITHUB_HOME", "/")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Invalid directory must cause panic")
		}
		defer os.Setenv("DOT_GITHUB_HOME", "")
	}()
	TemplateDir()
}

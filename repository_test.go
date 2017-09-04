package main

import (
	"strings"
	"testing"
)

func TestWebURL(t *testing.T) {
	r := NewRepositoryFromURL("https://github.com/rhysd/dot-github.git")
	if r.User != "rhysd" {
		t.Fatalf("User name is invalid: %v", r.User)
	}
	if r.Name != "dot-github" {
		t.Fatalf("Repository name is invalid: %v", r.Name)
	}
	if !strings.HasSuffix(r.Path, "dot-github") {
		t.Fatalf("Repository root must end with its name: %v", r.Path)
	}
}

func TestGitURL(t *testing.T) {
	r := NewRepositoryFromURL("git://github.com/rhysd/dot-github.git")
	if r.User != "rhysd" {
		t.Fatalf("User name is invalid: %v", r.User)
	}
	if r.Name != "dot-github" {
		t.Fatalf("Repository name is invalid: %v", r.Name)
	}
	if !strings.HasSuffix(r.Path, "dot-github") {
		t.Fatalf("Repository root must end with its name: %v", r.Path)
	}
}

func TestSshURL(t *testing.T) {
	r := NewRepositoryFromURL("git@github.com:rhysd/dot-github.git")
	if r.User != "rhysd" {
		t.Fatalf("User name is invalid: %v", r.User)
	}
	if r.Name != "dot-github" {
		t.Fatalf("Repository name is invalid: %v", r.Name)
	}
	if !strings.HasSuffix(r.Path, "dot-github") {
		t.Fatalf("Repository root must end with its name: %v", r.Path)
	}
}

func TestInvalidPath(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("URL without path must cause panic")
		}
	}()
	NewRepositoryFromURL("https://github.com")
}

func TestInvalidGitURL(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Invalid URL must cause panic")
		}
	}()
	NewRepositoryFromURL("invalid-blah")
}

func TestInvalidURLScheme(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Invalid Scheme must cause panic")
		}
	}()
	NewRepositoryFromURL("file://blah.jp")
}

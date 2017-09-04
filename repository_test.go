package main

import (
	"net/url"
	"strings"
	"testing"
)

func TestWebURL(t *testing.T) {
	u, _ := url.Parse("https://github.com/rhysd/dot-github.git")
	r := NewRepositoryFromURL(u)
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
	u, _ := url.Parse("git://github.com/rhysd/dot-github.git")
	r := NewRepositoryFromURL(u)
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
	u, err := url.Parse("git@github.com:rhysd/dot-github.git")
	if err != nil {
		t.Fatal(err)
	}
	r := NewRepositoryFromURL(u)
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
	u, _ := url.Parse("https://github.com")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("URL without path must cause panic")
		}
	}()
	NewRepositoryFromURL(u)
}

func TestInvalidGitURL(t *testing.T) {
	u, _ := url.Parse("invalid-blah")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Invalid URL must cause panic")
		}
	}()
	NewRepositoryFromURL(u)
}

func TestInvalidURLScheme(t *testing.T) {
	u, _ := url.Parse("file://blah.jp")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Invalid Scheme must cause panic")
		}
	}()
	NewRepositoryFromURL(u)
}

package main

import (
	"net/url"
	"strings"
	"testing"
)

func TestHttpsURL(t *testing.T) {
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
	u, _ := url.Parse("git@github.com:rhysd/dot-github.git")
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

package main

import (
	"path"
	"strings"
	"testing"
)

func TestRemoteURL(t *testing.T) {
	o := RemoteURL("origin")
	if o == nil {
		t.Fatalf("RemoteURL() must return non-empty string")
	}
	if !strings.Contains(o.String(), "dot-github") {
		t.Fatalf("'origin' remote url is invalid: %v", o)
	}
}

func TestGitRoot(t *testing.T) {
	r := GitRoot()
	if len(r) == 0 {
		t.Fatalf("GitRoot() must return non-empty string")
	}
	if !path.IsAbs(r) {
		t.Fatalf("GitRoot() must return absolute path but actually: %v", r)
	}
	if !strings.Contains(r, "dot-github") {
		t.Fatalf("GitRoot() must return path to its repository but actually: %v", r)
	}
}

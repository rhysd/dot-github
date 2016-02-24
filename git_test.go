package main

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"testing"
)

func TestRemoteURL(t *testing.T) {
	o := RemoteURL("origin")
	if o == nil {
		t.Fatalf("RemoteURL() must return non-empty URL")
	}
	if !strings.Contains(o.String(), "dot-github") {
		t.Fatalf("'origin' remote url is invalid: %v", o)
	}
}

func TestRemoteURLWithEnvVar(t *testing.T) {
	os.Setenv("DOT_GITHUB_GIT_CMD", "git")
	o := RemoteURL("origin")
	if o == nil {
		t.Fatalf("RemoteURL() must return non-empty URL")
	}
	if !strings.Contains(o.String(), "dot-github") {
		t.Fatalf("'origin' remote url is invalid: %v", o)
	}
	os.Setenv("DOT_GITHUB_GIT_CMD", "")
}

func TestRemoteURLWithInvalidEnvVar(t *testing.T) {
	os.Setenv("DOT_GITHUB_GIT_CMD", "unknown-command")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Uknown command for Git must cause panic")
		}
		os.Setenv("DOT_GITHUB_GIT_CMD", "")
	}()
	RemoteURL("origin")
}

func TestInvalidRemote(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Uknown remote must cause panic")
		}
	}()
	RemoteURL("uknown-remote-name")
}

func TestGitRoot(t *testing.T) {
	r := GitRoot()
	if len(r) == 0 {
		t.Fatalf("GitRoot() must return non-empty string")
	}
	if !filepath.IsAbs(r) {
		t.Fatalf("GitRoot() must return absolute path but actually: %v", r)
	}
	if !strings.Contains(r, "dot-github") {
		t.Fatalf("GitRoot() must return path to its repository but actually: %v", r)
	}
}

func TestNonGitRepo(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	wd, _ := os.Getwd()
	os.Chdir(u.HomeDir)
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("panic must occur when current dir is not Git repo")
		}
		os.Chdir(wd)
	}()
	GitRoot()
}

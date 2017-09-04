package main

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGitHubRemoteURL(t *testing.T) {
	o := GitHubRemoteURL("origin")
	if !strings.Contains(o, "dot-github") {
		t.Fatalf("'origin' remote url is invalid: %v", o)
	}
}

func TestGitHubRemoteURLWithEnvVar(t *testing.T) {
	os.Setenv("DOT_GITHUB_GIT_CMD", "git")
	o := GitHubRemoteURL("origin")
	if !strings.Contains(o, "dot-github") {
		t.Fatalf("'origin' remote url is invalid: %v", o)
	}
	os.Setenv("DOT_GITHUB_GIT_CMD", "")
}

func TestGitHubRemoteURLWithInvalidEnvVar(t *testing.T) {
	os.Setenv("DOT_GITHUB_GIT_CMD", "unknown-command")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Uknown command for Git must cause panic")
		}
		os.Setenv("DOT_GITHUB_GIT_CMD", "")
	}()
	GitHubRemoteURL("origin")
}

func TestInvalidRemote(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Uknown remote must cause panic")
		}
	}()
	GitHubRemoteURL("uknown-remote-name")
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
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	wd, _ := os.Getwd()
	os.Chdir(home)
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("panic must occur when current dir is not Git repo")
		}
		os.Chdir(wd)
	}()
	GitRoot()
}

func TestValidateURLs(t *testing.T) {
	var b bool
	b = ValidateGitHubURL("http://github.com/rhysd/dot-github.git")
	if !b {
		t.Errorf("Correct HTTP GitHub URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("https://github.com/rhysd/dot-github.git")
	if !b {
		t.Errorf("Correct HTTPS GitHub URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("git://github.com/rhysd/dot-github.git")
	if !b {
		t.Errorf("Correct GIT GitHub URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("git@github.com:rhysd/dot-github.git")
	if !b {
		t.Errorf("Correct SSH GitHub URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("http://github.company.com/rhysd/dot-github.git")
	if !b {
		t.Errorf("Correct HTTP GHE URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("https://github.company.com/rhysd/dot-github.git")
	if !b {
		t.Errorf("Correct HTTPS GHE URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("git://github.company.com/rhysd/dot-github.git")
	if !b {
		t.Errorf("Correct GIT GHE URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("git@github.company.com:rhysd/dot-github.git")
	if !b {
		t.Errorf("Correct SSH GHE URL was detected as invalid URL")
	}
}

func TestInvalidateURLs(t *testing.T) {
	var b bool
	b = ValidateGitHubURL("http://example.com/rhysd/dot-example.git")
	if b {
		t.Errorf("Correct HTTP GitHub URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("https://example.com/rhysd/dot-example.git")
	if b {
		t.Errorf("Correct HTTPS GitHub URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("git://example.com/rhysd/dot-example.git")
	if b {
		t.Errorf("Correct GIT GitHub URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("git@example.com:rhysd/dot-example.git")
	if b {
		t.Errorf("Correct SSH GitHub URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("http://example.company.com/rhysd/dot-example.git")
	if b {
		t.Errorf("Correct HTTP GHE URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("https://example.company.com/rhysd/dot-example.git")
	if b {
		t.Errorf("Correct HTTPS GHE URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("git://example.company.com/rhysd/dot-example.git")
	if b {
		t.Errorf("Correct GIT GHE URL was detected as invalid URL")
	}
	b = ValidateGitHubURL("git@example.company.com:rhysd/dot-example.git")
	if b {
		t.Errorf("Correct SSH GHE URL was detected as invalid URL")
	}
}

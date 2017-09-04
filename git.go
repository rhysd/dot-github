package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func gitCmdPath() string {
	cmd := os.Getenv("DOT_GITHUB_GIT_CMD")
	if len(cmd) == 0 {
		cmd = "git" // Default
	}

	path, err := exec.LookPath(cmd)
	if err != nil {
		panic("'" + cmd + "' command not found.  Specify $DOT_GITHUB_GIT_CMD properly.")
	}
	return path
}

func ValidateGitHubURL(u string) bool {
	return regexp.MustCompile(`^(:?https|http|git)://github(:?\..+)?\.com`).MatchString(u) ||
		regexp.MustCompile(`^git@github(:?\..+)?\.com:`).MatchString(u)
}

func GitHubRemoteURL(name string) string {
	cmd := exec.Command(gitCmdPath(), "ls-remote", "--get-url", name)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	u := string(out[:])
	if !ValidateGitHubURL(u) {
		panic("Invalid GitHub remote: " + name)
	}
	return strings.TrimSpace(u)
}

func GitRoot() string {
	cmd := exec.Command(gitCmdPath(), "rev-parse", "--show-cdup")
	out, err := cmd.Output()
	if err != nil {
		panic("Current directory is not in git repository")
	}
	rel := strings.TrimSpace(string(out[:]))
	if len(rel) == 0 {
		// Note:
		// Passing empty string to Abs() causes panic() in Windows
		rel = "."
	}
	root, err := filepath.Abs(rel)
	if err != nil {
		panic(err.Error() + ": " + rel)
	}
	return root
}

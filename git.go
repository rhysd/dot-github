package main

import (
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
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

func validateURL(u string) bool {
	return strings.HasPrefix(u, "https://") ||
		strings.HasPrefix(u, "http://") ||
		strings.HasPrefix(u, "git@") ||
		strings.HasPrefix(u, "git://")
}

func RemoteURL(name string) *url.URL {
	cmd := exec.Command(gitCmdPath(), "ls-remote", "--get-url", name)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	u := string(out[:])
	if !validateURL(u) {
		panic("Invalid remote: " + name)
	}
	url, err := url.Parse(strings.TrimSpace(u))
	if err != nil {
		panic(err)
	}

	return url
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

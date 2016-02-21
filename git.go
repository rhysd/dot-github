package main

import (
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func gitCmdPath() string {
	specified := os.Getenv("DOT_GITHUB_GIT_CMD")
	if len(specified) != 0 {
		return specified
	}

	path, err := exec.LookPath("git")
	if err != nil {
		panic("'git' command not found.  Consider to specify $DOT_GITHUB_GIT_CMD manually.")
	}
	return path
}
func RemoteURL(name string) *url.URL {
	cmd := exec.Command(gitCmdPath(), "ls-remote", "--get-url", name)
	out, err := cmd.Output()
	if err != nil {
		panic("Remote '" + name + "' was not found")
	}
	url, err := url.Parse(strings.TrimSpace(string(out[:])))
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
	root, err := filepath.Abs(strings.TrimSpace(string(out[:])))
	if err != nil {
		panic(err)
	}
	return root
}

package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func exitWithUsage() {
	fmt.Fprintln(
		os.Stderr,
		`$ dot-github [flags] [{dir}]

  A CLI tool to generate GitHub files such as CONTRIBUTING.md,
  ISSUE_TEMPLATE.md and PULLfrom template file.

  GitHub Blog: https://github.com/blog/2111-issue-and-pull-request-templates
  More usage:  https://github.com/rhysd/dot-github#readme

Flags:`)
	flag.PrintDefaults()
	os.Exit(0)
}

func exitWithVersion() {
	fmt.Println("0.0.0")
	os.Exit(0)
}

type Flags struct {
	Help    bool
	Version bool
}

func parseCmdArgs() *Flags {
	var (
		help    bool
		version bool
	)

	flag.BoolVar(&help, "help", false, "Show this help")
	flag.BoolVar(&version, "version", false, "Show version")
	flag.Parse()

	return &Flags{
		help,
		version,
	}
}

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

func originURL() *url.URL {
	cmd := exec.Command(gitCmdPath(), "ls-remote", "--get-url", "origin")
	out, err := cmd.Output()
	if err != nil {
		panic("'git ls-remote --git-url origin' failed: " + err.Error())
	}
	url, err := url.Parse(strings.TrimSpace(string(out[:])))
	if err != nil {
		panic(err.Error())
	}

	return url
}

type Repository struct {
	user string
	name string
}

func NewRepositoryFromHttpsURL(u *url.URL) *Repository {
	// TODO Check valid GitHub or GHE url
	if u.Path == "" {
		panic("Invalid https URL for GitHub: " + u.String())
	}
	split := strings.SplitN(u.Path[1:], "/", 2)
	return &Repository{
		split[0],
		strings.TrimSuffix(split[1], ".git"),
	}
}

func NewRepositoryFromGitURL(u *url.URL) *Repository {
	if !strings.HasPrefix(u.Path, "git@") || !strings.Contains(u.Path, ":") {
		panic("Invalid git@ URL for GitHub: " + u.String())
	}
	// TODO Check valid GitHub or GHE url
	split := strings.SplitN(
		strings.SplitN(u.Path, ":", 2)[1],
		"/",
		2,
	)
	return &Repository{
		split[0],
		strings.TrimSuffix(split[1], ".git"),
	}
}

func NewRepositoryFromURL(u *url.URL) *Repository {
	if u.Scheme == "https" {
		return NewRepositoryFromHttpsURL(u)
	} else if u.Scheme == "" {
		return NewRepositoryFromGitURL(u)
	}
	return nil
}

func main() {
	flags := parseCmdArgs()
	if flags.Help {
		exitWithUsage()
	} else if flags.Version {
		exitWithVersion()
	}

	// TODO
	fmt.Println(NewRepositoryFromURL(originURL()))
}

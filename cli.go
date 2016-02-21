package main

import (
	"flag"
	"fmt"
	"os"
)

func ExitWithUsage() {
	fmt.Fprintln(
		os.Stderr,
		`Usage: dot-github [flags]

  A CLI tool to generate GitHub files such as CONTRIBUTING.md,
  ISSUE_TEMPLATE.md and PULL_REQUEST_TEMPLATE.md from template files in
  '~/.github' directory.

  You can control which template should be used and it attempts to generate
  all by default.

  Below templates are looked by dot-github command.
  ($DOT_GITHUB_HOME is defaulted to $HOME)

    $DOT_GITHUB_HOME/.github/ISSUE_TEMPLATE.md
    $DOT_GITHUB_HOME/.github/PULL_REQUEST_TEMPLATE.md
    $DOT_GITHUB_HOME/.github/CONTRIBUTING.md

  You can use Golang's standard template for the template files.
  Below variables are available by default.

    .IsIssue        : (boolean) true when used for issue template
    .IsPullRequest  : (boolean) true when used for pull request template
    .IsContributing : (boolean) true when used for contributing template
    .RepoName       : (string) repository name
    .RepoUser       : (string) repository owner name

References:

  GitHub Blog:     https://github.com/blog/2111-issue-and-pull-request-templates
  More usage:      https://github.com/rhysd/dot-github#readme
  Golang template: https://golang.org/pkg/text/template/

Flags:`)
	flag.PrintDefaults()
	os.Exit(0)
}

func ExitWithVersion() {
	fmt.Println("1.0.0")
	os.Exit(0)
}

type Flags struct {
	Help             bool
	Version          bool
	IssueOnly        bool
	PROnly           bool
	ContributingOnly bool
}

func ParseCmdArgs() *Flags {
	var (
		help         bool
		version      bool
		issue        bool
		pr           bool
		contributing bool
	)

	flag.BoolVar(&help, "help", false, "Show this help")
	flag.BoolVar(&version, "version", false, "Show version")
	flag.BoolVar(&issue, "issue", false, "Import ISSUE_TEMPLATE.md only")
	flag.BoolVar(&pr, "pullrequest", false, "Import PULL_REQUEST_TEMPLATE.md only")
	flag.BoolVar(&contributing, "contributing", false, "Import CONTRIBUTING.md only")
	flag.Parse()

	return &Flags{
		help,
		version,
		issue,
		pr,
		contributing,
	}
}

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Parsed struct {
	Help             bool
	Version          bool
	IssueOnly        bool
	PROnly           bool
	ContributingOnly bool
	SelfUpdate       bool
	flags            *flag.FlagSet
}

const Usage = `Usage: dot-github [flags]

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

Flags:`

func (p *Parsed) ShowUsage(out io.Writer) {
	fmt.Fprintln(out, Usage)
	p.flags.SetOutput(out)
	p.flags.PrintDefaults()
}

func ParseCmdArgs(err_out io.Writer) (*Parsed, error) {
	var (
		help         bool
		version      bool
		issue        bool
		pr           bool
		contributing bool
		selfUpdate   bool
	)

	flags := flag.NewFlagSet("dot-github", flag.ContinueOnError)
	flags.SetOutput(err_out)

	flags.BoolVar(&help, "help", false, "Show this help")
	flags.BoolVar(&version, "version", false, "Show version")
	flags.BoolVar(&issue, "issue", false, "Import ISSUE_TEMPLATE.md only")
	flags.BoolVar(&pr, "pullrequest", false, "Import PULL_REQUEST_TEMPLATE.md only")
	flags.BoolVar(&contributing, "contributing", false, "Import CONTRIBUTING.md only")
	flags.BoolVar(&selfUpdate, "selfupdate", false, "Update dot-github binary itself")

	flags.Usage = func() {
		fmt.Fprintln(err_out, Usage)
		flags.PrintDefaults()
	}

	if err := flags.Parse(os.Args[1:]); err != nil {
		return nil, err
	}

	return &Parsed{
		help,
		version,
		issue,
		pr,
		contributing,
		selfUpdate,
		flags,
	}, nil
}

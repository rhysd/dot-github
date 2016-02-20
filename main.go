package main

import (
	"flag"
	"fmt"
	"os"
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

func main() {
	flags := parseCmdArgs()
	if flags.Help {
		exitWithUsage()
	} else if flags.Version {
		exitWithVersion()
	}
}

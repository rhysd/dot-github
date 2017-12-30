package main

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"io"
)

const version = "1.2.0"

type App struct {
	stdout io.Writer
	stderr io.Writer
}

func (a *App) selfUpdate() int {
	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, "rhysd/dot-github")
	if err != nil {
		fmt.Fprintln(a.stderr, err)
		return 1
	}

	if v.Equals(latest.Version) {
		fmt.Println("Current binary is the latest version", v)
	} else {
		fmt.Println("Successfully updated to version", latest.Version)
		fmt.Println("Release note:\n", latest.ReleaseNotes)
	}
	return 0
}

func (a *App) Run() int {
	parsed, err := ParseCmdArgs(a.stderr)
	if err != nil {
		return 1
	}
	if parsed.Help {
		parsed.ShowUsage(a.stderr)
		return 0
	} else if parsed.Version {
		fmt.Fprintln(a.stdout, version)
		return 0
	} else if parsed.SelfUpdate {
		return a.selfUpdate()
	}

	g := NewGenerator(
		TemplateDir(),
		NewRepositoryFromURL(GitHubRemoteURL("origin")),
	)
	if parsed.IssueOnly {
		g.GenerateIssueTemplate()
	}
	if parsed.PROnly {
		g.GeneratePRTemplate()
	}
	if parsed.ContributingOnly {
		g.GenerateContributingTemplate()
	}
	if !parsed.IssueOnly && !parsed.PROnly && !parsed.ContributingOnly {
		g.GenerateAllTemplates()
	}
	if !g.FileCreated {
		fmt.Fprintln(a.stdout, "No file created. Add template files to "+g.templateDir)
	}
	return 0
}

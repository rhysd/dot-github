package main

import (
	"fmt"
	"io"
)

type App struct {
	stdout io.Writer
	stderr io.Writer
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
		parsed.ShowVersion(a.stdout)
		return 0
	}

	g := NewGenerator(
		TemplateDir(),
		NewRepositoryFromURL(RemoteURL("origin")),
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

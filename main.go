package main

import (
	"fmt"
)

func main() {
	flags := ParseCmdArgs()
	if flags.Help {
		ExitWithUsage()
	} else if flags.Version {
		ExitWithVersion()
	}

	g := NewGenerator(
		TemplateDir(),
		NewRepositoryFromURL(RemoteURL("origin")),
	)
	if flags.IssueOnly {
		g.GenerateIssueTemplate()
	}
	if flags.PROnly {
		g.GeneratePRTemplate()
	}
	if flags.ContributingOnly {
		g.GenerateContributingTemplate()
	}
	if !flags.IssueOnly && !flags.PROnly && !flags.ContributingOnly {
		g.GenerateAllTemplates()
	}
	if !g.fileCreated {
		fmt.Println("No file created. Add template files to " + g.templateDir)
	}
}

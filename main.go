package main

import (
	"fmt"
	"os"
)

func main() {
	parsed, err := ParseCmdArgs(os.Stderr)
	if err != nil {
		os.Exit(1)
	}
	if parsed.Help {
		parsed.ShowUsage(os.Stderr)
		os.Exit(0)
	} else if parsed.Version {
		parsed.ShowVersion(os.Stdout)
		os.Exit(0)
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
	if !g.fileCreated {
		fmt.Println("No file created. Add template files to " + g.templateDir)
	}
}

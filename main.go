package main

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
	} else if flags.PROnly {
		g.GeneratePRTemplate()
	} else if flags.ContributingOnly {
		g.GenerateContributingTemplate()
	} else {
		g.GenerateAllTemplates()
	}
}

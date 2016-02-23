package main

import (
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"
)

func createAndChdirTo(dir string, issue_and_pr string, contributing string) {
	if err := os.MkdirAll(path.Join(dir, ".github"), os.ModeDir|os.ModePerm); err != nil {
		panic("Failed to make directory for testing")
	}
	if err := os.MkdirAll(path.Join(dir, "repo"), os.ModeDir|os.ModePerm); err != nil {
		panic("Failed to make repo directory for testing")
	}
	os.Chdir(dir)

	f1, err := os.Create(".github/ISSUE_AND_PULL_REQUEST_TEMPLATE.md")
	if err != nil {
		panic("Failed to make template file for issue and pull request")
	}
	defer f1.Close()
	f1.WriteString(issue_and_pr)

	f2, err := os.Create(path.Join(".github", "CONTRIBUTING.md"))
	if err != nil {
		panic("Failed to make template file for contributing")
	}
	defer f2.Close()
	f2.WriteString(contributing)
}

func chdirBackAndSweep(dir string) {
	os.Chdir("..")
	if err := os.RemoveAll(dir); err != nil {
		panic(err)
	}
}

func newGeneratorForTest(workdir string) *Generator {
	u, err := url.Parse("https://github.com/rhysd/dot-github.git")
	if err != nil {
		panic(err)
	}
	r := NewRepositoryFromURL(u)
	r.Path = path.Join(workdir, "repo")
	return NewGenerator(path.Join(workdir, ".github"), r)
}

func TestDotGithubDir(t *testing.T) {
	u, _ := url.Parse("https://github.com/rhysd/dot-github.git")
	g := NewGenerator(TemplateDir(), NewRepositoryFromURL(u))
	if !strings.HasSuffix(g.dotGithubDir, "dot-github/.github") {
		t.Fatalf("g.dotGithubDir must point to repository local .github directory: %v", g.dotGithubDir)
	}
}

func TestGeneratingAllTemplates(t *testing.T) {
	d := "test-generate-all-templates"
	createAndChdirTo(d, "### This is test for {{.RepoUser}}/{{.RepoName}}\n{{if .IsPullRequest}}pull request!{{else if .IsIssue}}issue!{{else if .IsContributing}}contributing!{{end}}", "This is contributing guide for {{.RepoUser}}/{{.RepoName}}")
	defer chdirBackAndSweep(d)

	w, _ := os.Getwd()
	g := newGeneratorForTest(w)
	g.GenerateAllTemplates()

	var (
		content string
		bytes   []byte
		err     error
	)

	bytes, err = ioutil.ReadFile(path.Join(w, "repo", ".github", "ISSUE_TEMPLATE.md"))
	if err != nil {
		t.Fatalf("ISSUE_TEMPLATE.md was not created")
	}
	content = string(bytes[:])
	if content != "### This is test for rhysd/dot-github\nissue!" {
		t.Errorf("ISSUE_TEMPLATE.md is invalid: %v", content)
	}

	bytes, err = ioutil.ReadFile(path.Join(w, "repo", ".github", "PULL_REQUEST_TEMPLATE.md"))
	if err != nil {
		t.Fatalf("PULL_REQUEST_TEMPLATE.md was not created")
	}
	content = string(bytes[:])
	if content != "### This is test for rhysd/dot-github\npull request!" {
		t.Errorf("PULL_REQUEST_TEMPLATE.md is invalid: %v", content)
	}

	bytes, err = ioutil.ReadFile(path.Join(w, "repo", ".github", "CONTRIBUTING.md"))
	if err != nil {
		t.Fatalf("CONTRIBUTING.md was not created")
	}
	content = string(bytes[:])
	if content != "This is contributing guide for rhysd/dot-github" {
		t.Errorf("CONTRIBUTING.md is invalid: %v", content)
	}

	if !g.FileCreated {
		t.Errorf("FileCreated flag was not set")
	}
}

func TestIgnoreNotExsistingTemplates(t *testing.T) {
	d := "test-not-existing"
	createAndChdirTo(d, "### This is test for {{.RepoUser}}/{{.RepoName}}\n{{if .IsPullRequest}}pull request!{{else if .IsIssue}}issue!{{else if .IsContributing}}contributing!{{end}}", "This is contributing guide for {{.RepoUser}}/{{.RepoName}}")
	defer chdirBackAndSweep(d)

	w, _ := os.Getwd()
	g := newGeneratorForTest(w)

	g.templateDir = path.Join(w, "unknown", ".github")

	// Template dir does not exist so it generates nothing
	g.GenerateAllTemplates()

	if g.FileCreated {
		t.Errorf("File is generated although template dir does not exist")
	}
}

func TestFailToParseTemplate(t *testing.T) {
	d := "test-not-existing"
	createAndChdirTo(d, "This causes error: {{", "")
	defer chdirBackAndSweep(d)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Parse error did not occur")
		}
	}()

	w, _ := os.Getwd()
	g := newGeneratorForTest(w)
	g.GenerateIssueTemplate()
}

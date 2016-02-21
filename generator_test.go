package main

import (
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"
)

// Create fixtures
func setup() {
	if err := os.MkdirAll(path.Join("dir-for-test", ".github"), os.ModeDir|os.ModePerm); err != nil {
		panic("Failed to make directory for testing")
	}
	if err := os.MkdirAll(path.Join("dir-for-test", "repo"), os.ModeDir|os.ModePerm); err != nil {
		panic("Failed to make repo directory for testing")
	}
	os.Chdir("dir-for-test")

	f1, err := os.Create(".github/ISSUE_AND_PULL_REQUEST_TEMPLATE.md")
	if err != nil {
		panic("Failed to make template file for issue and pull request")
	}
	defer f1.Close()
	f1.WriteString("### This is test for {{.RepoUser}}/{{.RepoName}}\n{{if .IsPullRequest}}pull request!{{else if .IsIssue}}issue!{{else if .IsContributing}}contributing!{{end}}")

	f2, err := os.Create(path.Join(".github", "CONTRIBUTING.md"))
	if err != nil {
		panic("Failed to make template file for contributing")
	}
	defer f2.Close()
	f2.WriteString("This is contributing guide for {{.RepoUser}}/{{.RepoName}}")
}

func teardown() {
	os.Chdir("..")
	if err := os.RemoveAll("dir-for-test"); err != nil {
		panic(err)
	}
}

func TestDotGithubDir(t *testing.T) {
	u, _ := url.Parse("https://github.com/rhysd/dot-github.git")
	g := NewGenerator(TemplateDir(), NewRepositoryFromURL(u))
	if !strings.HasSuffix(g.dotGithubDir, "dot-github/.github") {
		t.Fatalf("g.dotGithubDir must point to repository local .github directory: %v", g.dotGithubDir)
	}
}

func TestGeneratingFile(t *testing.T) {
	u, _ := url.Parse("https://github.com/rhysd/dot-github.git")
	r := NewRepositoryFromURL(u)
	w, _ := os.Getwd()
	r.Path = path.Join(w, "repo")
	g := NewGenerator(path.Join(w, ".github"), r)
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
		t.Fatalf("ISSUE_TEMPLATE.md is invalid: %v", content)
	}

	bytes, err = ioutil.ReadFile(path.Join(w, "repo", ".github", "PULL_REQUEST_TEMPLATE.md"))
	if err != nil {
		t.Fatalf("PULL_REQUEST_TEMPLATE.md was not created")
	}
	content = string(bytes[:])
	if content != "### This is test for rhysd/dot-github\npull request!" {
		t.Fatalf("PULL_REQUEST_TEMPLATE.md is invalid: %v", content)
	}

	bytes, err = ioutil.ReadFile(path.Join(w, "repo", ".github", "CONTRIBUTING.md"))
	if err != nil {
		t.Fatalf("CONTRIBUTING.md was not created")
	}
	content = string(bytes[:])
	if content != "This is contributing guide for rhysd/dot-github" {
		t.Fatalf("CONTRIBUTING.md is invalid: %v", content)
	}
}

func TestMain(m *testing.M) {
	setup()
	c := m.Run()
	teardown()
	os.Exit(c)
}

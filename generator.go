package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

type Generator struct {
	templateDir  string
	dotGithubDir string
	repo         *Repository
	FileCreated  bool
}

func NewGenerator(temp string, repo *Repository) *Generator {
	dotdir := path.Join(repo.Path, ".github")
	if _, err := os.Stat(dotdir); os.IsNotExist(err) {
		if err := os.MkdirAll(dotdir, os.ModeDir|os.ModePerm); err != nil {
			panic(err)
		}
	}
	return &Generator{
		temp,
		dotdir,
		repo,
		false,
	}
}

type Placeholders struct {
	IsPullRequest  bool
	IsIssue        bool
	IsContributing bool
	RepoUser       string
	RepoName       string
}

func (g *Generator) applyTemplate(src_path string, dst_path string) {
	dst, err := os.Create(dst_path)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	tmpl, err := template.ParseFiles(src_path)
	if err != nil {
		panic(err)
	}

	holders := Placeholders{
		strings.Contains(dst_path, "PULL_REQUEST_TEMPLATE.md"),
		strings.Contains(dst_path, "ISSUE_TEMPLATE.md"),
		strings.Contains(dst_path, "CONTRIBUTING.md"),
		g.repo.User,
		g.repo.Name,
	}

	if err := tmpl.Execute(dst, holders); err != nil {
		panic(err)
	}

	fmt.Println("Created " + dst_path)
}

func (g *Generator) generateFile(name string, fallback string) {
	src := path.Join(g.templateDir, name)
	if len(fallback) != 0 {
		if _, err := os.Stat(src); os.IsNotExist(err) {
			src = path.Join(g.templateDir, fallback)
		}
	}
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return
	}
	dst := path.Join(g.dotGithubDir, name)
	g.applyTemplate(src, dst)
	g.FileCreated = true
}

func (g *Generator) GenerateIssueTemplate() {
	g.generateFile("ISSUE_TEMPLATE.md", "ISSUE_AND_PULL_REQUEST_TEMPLATE.md")
}

func (g *Generator) GeneratePRTemplate() {
	g.generateFile("PULL_REQUEST_TEMPLATE.md", "ISSUE_AND_PULL_REQUEST_TEMPLATE.md")
}

func (g *Generator) GenerateContributingTemplate() {
	g.generateFile("CONTRIBUTING.md", "")
}

func (g *Generator) GenerateAllTemplates() {
	g.GenerateIssueTemplate()
	g.GeneratePRTemplate()
	g.GenerateContributingTemplate()
}

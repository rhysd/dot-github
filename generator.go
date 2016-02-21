package main

import (
	"io"
	"os"
	"path"
)

type Generator struct {
	templateDir  string
	dotGithubDir string
	repo         *Repository
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
	}
}

func (g *Generator) applyTemplate(src_path string, dst_path string) {
	// XXX: Simply copy file
	src, err := os.Open(src_path)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Open(dst_path)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		panic(err)
	}
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
	dst := path.Join(g.repo.Path, name)
	g.applyTemplate(src, dst)
}

func (g *Generator) GenerateIssueTemplate() {
	g.generateFile("ISSUE_TEMPLATE.md", "TEMPLATE.md")
}

func (g *Generator) GeneratePRTemplate() {
	g.generateFile("PULL_REQUEST_TEMPLATE.md", "TEMPLATE.md")
}

func (g *Generator) GenerateContributingTemplate() {
	g.generateFile("CONTRIBUTING.md", "")
}

func (g *Generator) GenerateAllTemplates() {
	g.GenerateIssueTemplate()
	g.GeneratePRTemplate()
	g.GenerateContributingTemplate()
}

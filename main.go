package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

func exitWithUsage() {
	fmt.Fprintln(
		os.Stderr,
		`Usage: dot-github [flags]

  A CLI tool to generate GitHub files such as CONTRIBUTING.md,
  ISSUE_TEMPLATE.md and PR_TEMPLATE.md from template files in ~/.github
  directory.
  You can control which template should be used and it attempts to generate
  all by default.

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
	Help             bool
	Version          bool
	IssueOnly        bool
	PROnly           bool
	ContributingOnly bool
}

func parseCmdArgs() *Flags {
	var (
		help         bool
		version      bool
		issue        bool
		pr           bool
		contributing bool
	)

	flag.BoolVar(&help, "help", false, "Show this help")
	flag.BoolVar(&version, "version", false, "Show version")
	flag.BoolVar(&issue, "issue", false, "Import ISSUE_TEMPLATE.md only")
	flag.BoolVar(&pr, "pr", false, "Import PR_TEMPLATE.md only")
	flag.BoolVar(&contributing, "contributing", false, "Import CONTRIBUTING.md only")
	flag.Parse()

	return &Flags{
		help,
		version,
		issue,
		pr,
		contributing,
	}
}

func gitCmdPath() string {
	specified := os.Getenv("DOT_GITHUB_GIT_CMD")
	if len(specified) != 0 {
		return specified
	}

	path, err := exec.LookPath("git")
	if err != nil {
		panic("'git' command not found.  Consider to specify $DOT_GITHUB_GIT_CMD manually.")
	}
	return path
}

func RemoteURL(name string) *url.URL {
	cmd := exec.Command(gitCmdPath(), "ls-remote", "--get-url", name)
	out, err := cmd.Output()
	if err != nil {
		panic("Remote '" + name + "' was not found")
	}
	url, err := url.Parse(strings.TrimSpace(string(out[:])))
	if err != nil {
		panic(err)
	}

	return url
}

type Repository struct {
	User string
	Name string
	Path string
}

func GitRoot() string {
	cmd := exec.Command(gitCmdPath(), "rev-parse", "--show-cdup")
	out, err := cmd.Output()
	if err != nil {
		panic("Current directory is not in git repository")
	}
	root, err := filepath.Abs(strings.TrimSpace(string(out[:])))
	if err != nil {
		panic(err)
	}
	return root
}

func NewRepositoryFromHttpsURL(u *url.URL) *Repository {
	// TODO Check valid GitHub or GHE url
	if u.Path == "" {
		panic("Invalid https URL for GitHub: " + u.String())
	}
	split := strings.SplitN(u.Path[1:], "/", 2)
	return &Repository{
		split[0],
		strings.TrimSuffix(split[1], ".git"),
		GitRoot(),
	}
}

func NewRepositoryFromGitURL(u *url.URL) *Repository {
	if !strings.HasPrefix(u.Path, "git@") || !strings.Contains(u.Path, ":") {
		panic("Invalid git@ URL for GitHub: " + u.String())
	}
	// TODO Check valid GitHub or GHE url
	split := strings.SplitN(
		strings.SplitN(u.Path, ":", 2)[1],
		"/",
		2,
	)
	return &Repository{
		split[0],
		strings.TrimSuffix(split[1], ".git"),
		GitRoot(),
	}
}

func NewRepositoryFromURL(u *url.URL) *Repository {
	if u.Scheme == "https" {
		return NewRepositoryFromHttpsURL(u)
	} else if u.Scheme == "" {
		return NewRepositoryFromGitURL(u)
	}
	return nil
}

func baseDir() string {
	env := os.Getenv("DOT_GITHUB_HOME")
	if len(env) != 0 {
		return env
	}

	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	return u.HomeDir
}

func TemplateDir() string {
	d := path.Join(baseDir(), ".github")
	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err := os.MkdirAll(d, os.ModeDir|0644); err != nil {
			panic(err)
		}
	}
	return d
}

type Importer struct {
	templateDir  string
	dotGithubDir string
	repo         *Repository
}

func NewImporter(temp string, repo *Repository) *Importer {
	dotdir := path.Join(repo.Path, ".github")
	if _, err := os.Stat(dotdir); os.IsNotExist(err) {
		if err := os.MkdirAll(dotdir, os.ModeDir|0644); err != nil {
			panic(err)
		}
	}
	return &Importer{
		temp,
		dotdir,
		repo,
	}
}

func (i *Importer) applyTemplate(src_path string, dst_path string) {
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

func (i *Importer) importFile(name string, fallback string) {
	src := path.Join(i.templateDir, name)
	if len(fallback) != 0 {
		if _, err := os.Stat(src); os.IsNotExist(err) {
			src = path.Join(i.templateDir, fallback)
		}
	}
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return
	}
	dst := path.Join(i.repo.Path, name)
	i.applyTemplate(src, dst)
}

func (i *Importer) ImportIssueTemplate() {
	i.importFile("ISSUE_TEMPLATE.md", "TEMPLATE.md")
}

func (i *Importer) ImportPRTemplate() {
	i.importFile("PR_TEMPLATE.md", "TEMPLATE.md")
}

func (i *Importer) ImportContributingTemplate() {
	i.importFile("CONTRIBUTING.md", "")
}

func (i *Importer) ImportAllTemplates() {
	i.ImportIssueTemplate()
	i.ImportPRTemplate()
	i.ImportContributingTemplate()
}

func main() {
	flags := parseCmdArgs()
	if flags.Help {
		exitWithUsage()
	} else if flags.Version {
		exitWithVersion()
	}

	i := NewImporter(
		TemplateDir(),
		NewRepositoryFromURL(RemoteURL("origin")),
	)
	if flags.IssueOnly {
		i.ImportIssueTemplate()
	} else if flags.PROnly {
		i.ImportPRTemplate()
	} else if flags.ContributingOnly {
		i.ImportContributingTemplate()
	} else {
		i.ImportAllTemplates()
	}
}

`.github` Directory Generator for Your Repository
=================================================
[![Build Status](https://travis-ci.org/rhysd/dot-github.svg?branch=master)](https://travis-ci.org/rhysd/dot-github)

GitHub now supports [issue and pull request template](https://github.com/blog/2111-issue-and-pull-request-templates).  This repository provides `dot-github` command to generate the template files automatically for your GitHub repositories.

## Getting Started

### 1. Installation

```sh
$ go get github.com/rhysd/dot-github
```

### 2. Write Your Template Files

```sh
$ mkdir -p ~/.github && cd ~/.github
$ $EDITOR ISSUE_AND_PULL_REQUEST.md
$ $EDITOR CONTRIBUTING.md
```

Please read below instruction detail about template file

**Note:** You can change the home directory for `dot-github` by `$DOT_GITHUB_HOME` environment variable.

### 3. Generate `.github`

```sh
$ cd your-repo
$ dot-github
$ git add .github
```

### 4. Tweak Generated Files

Tweak generated files in `your-repo/.github/*` for your repository-specific information.

## Writing Template File

You can see [example direcotry in this repository](exapmle/) for real world examples.

`dot-github` looks below template files

| File Path                                                     | Description                                                                               |
| ------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `$DOT_GITHUB_HOME/.github/ISSUE_TEMPLATE.md`                  | Template for issues.                                                                      |
| `$DOT_GITHUB_HOME/.github/PULL_REQUEST_TEMPLATE.md`           | Template for pull requests.                                                               |
| `$DOT_GITHUB_HOME/.github/ISSUE_AND_PULL_REQUEST_TEMPLATE.md` | If above files are not found, this file is used for template of issues and pull requests. |
| `$DOT_GITHUB_HOME/.github/CONTRIBUTING.md`                    | Template for contributing guideline.                                                      |

Note that `$DOT_GITHUB_HOME` is an environment variable.  You can specify your favorite directory to put template files.  Default directory for it is `~`.

Above template files are parsed as [Golang's standard text template](https://golang.org/pkg/text/template/).  Below variables are available in template.  They are useful to write flexible and common template files for each repositories.

| Variable Name     | Type      | Description                               |
| ----------------- | --------- | ----------------------------------------- |
| `.IsIssue`        | *boolean* | True when used for issue template.        |
| `.IsPullRequest`  | *boolean* | True when used for pull request template. |
| `.IsContributing` | *boolean* | True when used for contributing template. |
| `.RepoName`       | *string*  | Repository name.                          |
| `.RepoUser`       | *string*  | Repository owner name.                    |

## Template Examples

### Template files

- `~/.github/ISSUE_AND_PULL_REQUEST_TEMPLATE.md`

```
{{if .IsIssue}}
### Expected Behavior


### Actual Behavior


{{end}}
{{if .IsPullRequest}}
### Fix or Enhancement?


- [ ] All tests passed
{{end}}

### Environment
- OS: Write here
- Go version: Write here
```

- `~/.github/CONTRIBUTING.md`

```
Thank you for contributing {{.RepoName}}!
=========================================

Please follow issue/PR template.
```

### Generated Files

- `/path/to/your-repo/.github/ISSUE_TEMPLATE.md`

```
### Expected Behavior


### Actual Behavior


### Environment
- OS: Write here
- Go version: Write here
```

- `/path/to/your-repo/.github/PULL_REQUEST_TEMPLATE.md`

```
### Fix or Enhancement?


- [ ] All tests passed

### Environment
- OS: Write here
- Go version: Write here
```

- `/path/to/your-repo/.github/CONTRIBUTING.md`

```
Thank you for contributing my-project!
=========================================

Please follow issue/PR template.
```

## References

- [Setting guidelines for repository contributors](https://help.github.com/articles/setting-guidelines-for-repository-contributors/)
- [Creating an issue template for your repository](https://help.github.com/articles/creating-an-issue-template-for-your-repository/)
- [Creating a pull request template for your repository](https://help.github.com/articles/creating-a-pull-request-template-for-your-repository/)

## License

This software is distributed under [MIT license](LICENSE.txt).

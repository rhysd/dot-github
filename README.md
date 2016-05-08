`.github` Directory Generator
=============================

| Linux and OS X | Windows | Coverage |
| -------------- | ------- | -------- |
| [![Build Status](https://travis-ci.org/rhysd/dot-github.svg?branch=master)](https://travis-ci.org/rhysd/dot-github) | [![Build status](https://ci.appveyor.com/api/projects/status/bjat5jyqmcgjfvwd?svg=true)](https://ci.appveyor.com/project/rhysd/dot-github) | [![Coverage Status](https://coveralls.io/repos/github/rhysd/dot-github/badge.svg?branch=master)](https://coveralls.io/github/rhysd/dot-github?branch=master) |

GitHub now supports [issue and pull request template](https://github.com/blog/2111-issue-and-pull-request-templates).  This repository provides `dot-github` command to generate the template files automatically for your GitHub repositories.  This also enables to manage template files in dotfiles for all of your machines.

![screenshot](https://raw.githubusercontent.com/rhysd/ss/master/dot-github/main.gif)

## Getting Started

### 1. Installation

`go get` command

```sh
$ go get github.com/rhysd/dot-github
```

or [released binaries](https://github.com/rhysd/dot-github/releases)

```sh
cd /path/to/Downloads  # Download binary for your platform
chmod +x dot-github_your_platform
mv dot-github_your_platform /usr/local/bin/dot-github
```

### 2. Write Your Template Files

```sh
$ mkdir -p ~/.github && cd ~/.github
$ $EDITOR ISSUE_AND_PULL_REQUEST_TEMPLATE.md
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

Tweak generated files in `your-repo/.github/*` for your project-specific information.

## Writing Template File

You can see [example directory in this repository](example/) for real world examples.

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
- [Issue と PR のテンプレートジェネレータつくった (Japanese Blog Post)](http://rhysd.hatenablog.com/entry/2016/02/21/233643)

## License

This software is distributed under [MIT license](LICENSE.txt).

package main

import (
	"net/url"
	"strings"
)

type Repository struct {
	User string
	Name string
	Path string
}

func NewRepositoryFromWebURL(u *url.URL) *Repository {
	// TODO Check valid GitHub or GHE url
	if u.Path == "" {
		panic("Invalid https URL for GitHub repository: " + u.String())
	}
	split := strings.SplitN(u.Path[1:], "/", 2)
	return &Repository{
		split[0],
		strings.TrimSuffix(split[1], ".git"),
		GitRoot(),
	}
}

func NewRepositoryFromSshURL(u *url.URL) *Repository {
	if !strings.HasPrefix(u.Path, "git@") || !strings.Contains(u.Path, ":") {
		panic("Invalid git@ URL for GitHub repository: " + u.String())
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
	switch u.Scheme {
	case "https":
		return NewRepositoryFromWebURL(u)
	case "git":
		return NewRepositoryFromWebURL(u)
	case "":
		return NewRepositoryFromSshURL(u)
	}
	panic("Invalid URL for GitHub: " + u.String())
}

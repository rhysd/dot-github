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

func NewRepositoryFromSshURL(u string) *Repository {
	if !strings.HasPrefix(u, "git@") || !strings.Contains(u, ":") {
		panic("Invalid git@ URL for GitHub repository: " + u)
	}
	// TODO Check valid GitHub or GHE url
	split := strings.SplitN(
		strings.SplitN(u, ":", 2)[1],
		"/",
		2,
	)
	return &Repository{
		split[0],
		strings.TrimSuffix(split[1], ".git"),
		GitRoot(),
	}
}

func NewRepositoryFromURL(s string) *Repository {
	u, err := url.Parse(s)
	if err != nil {
		return NewRepositoryFromSshURL(s)
	}
	switch u.Scheme {
	case "https":
		return NewRepositoryFromWebURL(u)
	case "git":
		return NewRepositoryFromWebURL(u)
	default:
		panic("Invalid URL for GitHub: " + s)
	}
}

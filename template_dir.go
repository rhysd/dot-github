package main

import (
	"os"
	"os/user"
	"path"
)

func baseDir() string {
	dir := os.Getenv("DOT_GITHUB_HOME")
	if len(dir) == 0 {
		u, err := user.Current()
		if err != nil {
			panic(err)
		}
		dir = u.HomeDir
	}
	return dir
}

func TemplateDir() string {
	d := path.Join(baseDir(), ".github")
	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err := os.MkdirAll(d, os.ModeDir|os.ModePerm); err != nil {
			panic(err)
		}
	}
	return d
}

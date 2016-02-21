package main

import (
	"os"
	"os/user"
	"path"
)

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

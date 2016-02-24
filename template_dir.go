package main

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
)

func baseDir() string {
	dir := os.Getenv("DOT_GITHUB_HOME")
	if len(dir) == 0 {
		var err error
		dir, err = homedir.Dir()
		if err != nil {
			panic(err)
		}
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

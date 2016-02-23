package main

import (
	"os"
)

func main() {
	app := &App{os.Stdout, os.Stderr}
	os.Exit(app.Run())
}

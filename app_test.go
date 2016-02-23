package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	os.Args = []string{"dot-github"}
	o := &bytes.Buffer{}
	e := &bytes.Buffer{}
	app := &App{o, e}
	app.Run()
	if e.String() != "" {
		t.Errorf("Some error was output while command running")
	}
}

func TestRunWithFlags(t *testing.T) {
	os.Args = []string{"dot-github", "-issue", "-pullrequest", "-contributing"}
	o := &bytes.Buffer{}
	e := &bytes.Buffer{}
	app := &App{o, e}
	app.Run()
	if e.String() != "" {
		t.Errorf("Some error was output while command running")
	}
}

func TestInvalidFlag(t *testing.T) {
	os.Args = []string{"dot-github", "-unknown"}
	o := &bytes.Buffer{}
	e := &bytes.Buffer{}
	app := &App{o, e}
	app.Run()
	if !strings.Contains(e.String(), "Usage:") {
		t.Errorf("Command must show usage on invalid flag")
	}
}

func TestHelpFlag(t *testing.T) {
	os.Args = []string{"dot-github", "-help"}
	o := &bytes.Buffer{}
	e := &bytes.Buffer{}
	app := &App{o, e}
	app.Run()
	if !strings.Contains(e.String(), "Usage:") {
		t.Errorf("Command must show usage on -help flag")
	}
}

func TestVersionFlag(t *testing.T) {
	os.Args = []string{"dot-github", "-version"}
	o := &bytes.Buffer{}
	e := &bytes.Buffer{}
	app := &App{o, e}
	app.Run()
	if !strings.Contains(o.String(), ".") {
		t.Errorf("Command must show version on -version flag")
	}

}

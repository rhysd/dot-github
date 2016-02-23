package main

import (
	"bytes"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestSpecifiedFlags(t *testing.T) {
	os.Args = []string{"dot-github", "-issue", "-help", "-version", "-pullrequest", "-contributing"}
	p, err := ParseCmdArgs(os.Stderr)
	if err != nil {
		t.Fatalf("must be parsed without error")
	}
	if !p.IssueOnly {
		t.Errorf("-issue must be looked")
	}
	if !p.Help {
		t.Errorf("-help must be looked")
	}
	if !p.PROnly {
		t.Errorf("-pullrequest must be looked")
	}
	if !p.ContributingOnly {
		t.Errorf("-contributing must be looked")
	}
	if !p.Version {
		t.Errorf("-version must be looked")
	}
}

func TestUnspecifiedFlags(t *testing.T) {
	os.Args = []string{"dot-github"}
	p, err := ParseCmdArgs(os.Stderr)
	if err != nil {
		t.Fatalf("must be parsed without error")
	}
	if p.IssueOnly {
		t.Errorf("-issue is not invalid default value")
	}
	if p.Help {
		t.Errorf("-help is not invalid default value")
	}
	if p.PROnly {
		t.Errorf("-pullrequest is not invalid default value")
	}
	if p.ContributingOnly {
		t.Errorf("-contributing is not invalid default value")
	}
	if p.Version {
		t.Errorf("-version is not invalid default value")
	}
}

func TestUndefinedFlags(t *testing.T) {
	os.Args = []string{"dot-github", "-unknown"}
	buf := &bytes.Buffer{}
	_, err := ParseCmdArgs(buf)
	if err == nil {
		t.Fatalf("ignores unknown flag")
	}
	if !strings.Contains(buf.String(), "Usage:") {
		t.Fatalf("does not output usage")
	}
}

func TestShowVersion(t *testing.T) {
	os.Args = []string{"dot-github"}
	p, _ := ParseCmdArgs(os.Stderr)
	buf := &bytes.Buffer{}
	p.ShowVersion(buf)
	if m, err := regexp.Match(`\d\.\d\.\d`, buf.Bytes()); !m || err != nil {
		t.Errorf("invalid version output: " + buf.String())
	}
}

func TestShowHelp(t *testing.T) {
	os.Args = []string{"dot-github"}
	p, _ := ParseCmdArgs(os.Stderr)
	buf := &bytes.Buffer{}
	p.ShowUsage(buf)
	if !strings.Contains(buf.String(), "Usage:") {
		t.Fatalf("invalid help output")
	}
}

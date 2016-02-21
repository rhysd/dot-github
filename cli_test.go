package main

import (
	"os"
	"testing"
)

func TestFlags(t *testing.T) {
	os.Args = []string{"dot-github", "-issue", "-help", "-version", "-pullrequest", "-contributing"}
	f := ParseCmdArgs()
	if !f.IssueOnly {
		t.Fatalf("-issue must be looked")
	}
	if !f.Help {
		t.Fatalf("-help must be looked")
	}
	if !f.PROnly {
		t.Fatalf("-pullrequest must be looked")
	}
	if !f.ContributingOnly {
		t.Fatalf("-contributing must be looked")
	}
	if !f.Version {
		t.Fatalf("-version must be looked")
	}
}

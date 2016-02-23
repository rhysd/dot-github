package main

import (
	"os"
	"testing"
)

func TestFlags(t *testing.T) {
	os.Args = []string{"dot-github", "-issue", "-help", "-version", "-pullrequest", "-contributing"}
	f := ParseCmdArgs()
	if !f.IssueOnly {
		t.Errorf("-issue must be looked")
	}
	if !f.Help {
		t.Errorf("-help must be looked")
	}
	if !f.PROnly {
		t.Errorf("-pullrequest must be looked")
	}
	if !f.ContributingOnly {
		t.Errorf("-contributing must be looked")
	}
	if !f.Version {
		t.Errorf("-version must be looked")
	}
}

func TestExtWithUsage(t *testing.T) {

}

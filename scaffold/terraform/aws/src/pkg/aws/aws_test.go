package aws

import (
	"os"
	"testing"
)

func TestRepoPath(t *testing.T) {
	input := Input{
		ProjectName: "test",
		ProjectPath: os.TempDir(),
	}

	want := os.TempDir() + "/test"
	got := input.RepoPath()
	if want != got {
		t.Errorf("RepoPath got %v, want %v", got, want)
	}
}

func TestRun(t *testing.T) {
	input := Input{
		ProjectName: "test",
		ProjectPath: os.TempDir(),
	}
	input.Run()

	if _, err := os.Stat(input.RepoPath()); os.IsNotExist(err) {
		t.Errorf("Run got %v, want %v", err, "project created")
	}
}

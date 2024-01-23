package tagtolibv_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ParfenovVS/mobile-devops/tagtolibv"
)

func TestGetCurrentBranch(t *testing.T) {
	defaultWd, _ := os.Getwd()
	defer os.Chdir(defaultWd)

	dir, err := createTempRepository(".temp_TestGetCurrentBranch")
	if err != nil {
		t.Fatalf("cannot create temp repository: %s", err.Error())
	}
	defer os.RemoveAll(dir)

	cmd := exec.Command("git", "checkout", "-b", "test_branch")
	cmd.Dir = dir
	cmd.Run()
	exp := "test_branch"

	branch, err := tagtolibv.GetCurrentBranch(dir)
	if err != nil {
		t.Errorf("failed with %q", err)
	} else if exp != branch {
		t.Errorf("expected %q, got %q instead.", exp, branch)
	}
}

func TestGetTags(t *testing.T) {
	defaultWd, _ := os.Getwd()
	defer os.Chdir(defaultWd)

	dir, err := createTempRepository(".temp_TestGetTags")
	if err != nil {
		t.Fatalf("cannot create temp repository: %s", err.Error())
	}
	defer os.RemoveAll(dir)

	expTags := []string{
		"v1.0.0",
		"v1.1.0-beta",
	}

	os.Create("temp")
	cmd := exec.Command("git", "add", "temp")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("cannot add to git: %s, %s", err.Error(), string(out))
	}
	cmd = exec.Command("git", "commit", "-m", "\"add temp\"")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("cannot commit: %s, %s", err.Error(), string(out))
	}

	for _, tag := range expTags {
		cmd := exec.Command("git", "tag", tag)
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("cannot create tag %s: %s, %s", tag, err.Error(), string(out))
		}
	}

	tags, err := tagtolibv.GetTags(dir)
	if err != nil {
		t.Fatalf("cannot get tags: %s", err.Error())
	}

	if !reflect.DeepEqual(expTags, tags) {
		t.Errorf("expected:\n")
		for _, tag := range expTags {
			t.Errorf("%q\n", tag)
		}
		t.Errorf("actual:\n")
		for _, tag := range tags {
			t.Errorf("%q\n", tag)
		}
	}
}

func createTempRepository(name string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot get work directory")
		return "", err
	}

	err = os.Mkdir(name, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot create temp directory")
		return "", err
	}

	repoPath := filepath.Join(wd, name)
	os.Chdir(repoPath)

	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "cannot create repository")
		return "", err
	}

	fmt.Printf("created repository: %s\n", repoPath)
	return repoPath, nil
}

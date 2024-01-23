package tagtolibv

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetCurrentBranch(path string) (string, error) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("cannot choose dir %s", path))
		return "", err
	}

	cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot get current branch")
		return "", err
	}

	return strings.Trim(string(out), "\n"), nil
}

func GetTags(path string) ([]string, error) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("cannot choose dir %s", path))
		return nil, err
	}
	cmd := exec.Command("git", "tag", "-l", "v*")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("cannot get tags: %s", err.Error()))
		return nil, err
	}

	tags := strings.Split(string(out), "\n")
	tags = tags[:len(tags)-1]

	return tags, nil
}

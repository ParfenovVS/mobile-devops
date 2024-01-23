package tagtolibv

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/BurntSushi/toml"
)

func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot get current branch")
		return "", err
	}

	return strings.Trim(string(out), "\n"), nil
}

func GetTags() ([]string, error) {
	cmd := exec.Command("git", "tag", "-l", "v*")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("cannot get tags: %s", err.Error()))
		return nil, err
	}

	tags := strings.Split(string(out), "\n")
	if len(tags) > 0 {
		tags = tags[:len(tags)-1]
	}

	return tags, nil
}

type tomlConfig struct {
	Versions map[string]string
}

func GetLibVersion(tag string, lib string) (string, error) {
	cmd := exec.Command("git", "checkout", tag)
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("cannot checkout tag %s: %s", tag, err.Error()))
		return "", nil
	}

	tomlFile := "gradle/libs.versions.toml"
	var config tomlConfig
	_, err := toml.DecodeFile(tomlFile, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", nil
	}

	version := config.Versions[lib]

	return version, nil
}

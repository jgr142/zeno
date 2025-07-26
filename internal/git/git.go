package git

import (
	"os/exec"
	"strings"
)

func Rebase() {

}

func CurBranch() (string, error) {
	cmd := exec.Command("git", "branch")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	outStr := string(out)
	starIdx := strings.IndexRune(outStr, '*')
	outStr = outStr[starIdx+2:] // Removes star + space after it
	nLineIdx := strings.IndexRune(outStr, '\n')
	outStr = outStr[:nLineIdx]

	return outStr, nil
}

func AddCommitPush(commitMessage string) error {
	branch, branchErr := CurBranch()
	if branchErr != nil {
		return branchErr
	}

	addCmd := exec.Command("git", "add", ".")
	_, addErr := addCmd.Output()
	if addErr != nil {
		return addErr
	}

	commitCmd := exec.Command("git", "commit", "-m", commitMessage)
	_, commitErr := commitCmd.Output()
	if commitErr != nil {
		return commitErr
	}

	pushCommand := exec.Command("git", "push", "origin", branch)
	_, pushErr := pushCommand.Output()
	if pushErr != nil {
		return pushErr
	}

	return nil
}

func Status() ([]byte, error) {
	cmd := exec.Command("git", "status")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func Log() ([]byte, error) {
	cmd := exec.Command("git", "log")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func Pull() error {
	cmd := exec.Command("git", "pull")
	_, err := cmd.Output()
	return err
}

func Stash() error {
	cmd := exec.Command("git", "stash")
	_, err := cmd.Output()
	return err
}

func Pop() error {
	cmd := exec.Command("git", "pop", "stash")
	_, err := cmd.Output()
	return err
}

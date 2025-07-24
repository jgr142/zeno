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

package git

import (
	"my-git/errors"
	"os/exec"
	"regexp"
	"strings"
)

// CheckoutNewBranch - does `git checkout -b {branch name}`
var CheckoutNewBranch = []string{"checkout", "-b"}

// CheckoutBranch - does `git checkout {branch name}`
var CheckoutBranch = []string{"checkout"}

// Branch - a struct with branch name and `current` flag.
type Branch struct {
	Name    string
	Currect bool
}

// Commit - a struct with commit name and hash.
type Commit struct {
	Name string
	Hash string
}

// Branches - get list of branches for current repository.
func Branches() []Branch {
	var results []Branch
	out, err := exec.Command("git", "branch", "-l").Output()
	errors.CheckError(err)
	re := regexp.MustCompile(`\s|\*`)

	for _, str := range strings.Split(string(out), "\n") {
		current, _ := regexp.MatchString(`\*`, str)
		branch := re.ReplaceAllString(str, "")
		if len(branch) == 0 {
			continue
		}
		results = append(results, Branch{Name: branch, Currect: current})
	}
	return results
}

// Commits - get list of commits.
func Commits() []Commit {
	var results []Commit
	out, err := exec.Command("git", "log", "--oneline").Output()
	errors.CheckError(err)

	for _, str := range strings.Split(string(out), "\n") {
		if len(str) == 0 {
			continue
		}
		tokens := strings.Split(str, " ")
		hash := tokens[0]
		name := strings.Join(tokens[1:], " ")
		results = append(results, Commit{Name: name, Hash: hash})
	}
	return results
}

// RunInlineCommand - run a simple git command, returning the output.
func RunInlineCommand(args ...string) (string, error) {
	out, err := exec.Command("git", args...).Output()
	return string(out), err
}

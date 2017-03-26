package main

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// GetBranches - get list of branches.
func GetBranches() []string {
	cmd := exec.Command("git", "branch")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	var processedLines []string
	lines := strings.Split(out.String(), "\n")
	r := regexp.MustCompile(`\*|\s+`)

	for _, line := range lines {
		processedLine := r.ReplaceAllString(line, "")
		if len(processedLine) > 0 {
			processedLines = append(processedLines, processedLine)
		}
	}
	return processedLines
}

// SetBranch - set branch as current.
func SetBranch(branch string) string {
	cmd := exec.Command("git", "checkout", branch)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/fatih/color"
)

func getBranchName(cwd string) (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()

	return strings.TrimSuffix(string(out), "\n"), err
}

func RenderInputLine() {
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	printPrefix(cwd)
}

func printPrefix(cwd string) {
	color.New(color.FgCyan).Add(color.Bold).Printf("%s ", path.Base(cwd))
	branch, err := getBranchName(cwd)
	if err == nil {
		color.New(color.FgBlue).Add(color.Bold).Print("(")
		color.New(color.FgRed).Add(color.Bold).Printf("%s", branch)
		color.New(color.FgBlue).Add(color.Bold).Print(") ")
	}
	color.New(color.FgCyan).Add(color.Bold).Print("> ")
}

package utils

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func ExecInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("cd requires path")
		}

		if err := os.Chdir(args[1]); err != nil {
			return errors.New("no such dir")
		}

		return nil
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

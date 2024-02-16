package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bohuslavsemenov/jsh/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		utils.RenderInputLine()

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err = utils.ExecInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

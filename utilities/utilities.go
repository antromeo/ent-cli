package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadString(inputText string, required bool) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(inputText + ": ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if required {
		for len(input) == 0 {
			fmt.Print("Input cannot be empty. Please enter a non-empty string: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
		}
	}

	return input
}

func IsEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}
	return false
}

func ShowAdditionalCommandsInHelp() {
	fmt.Println("ADDITIONAL COMMANDS")
	fmt.Println("  deploy       Generates the CR and deploys it to the currently attached EntandoApp")
	fmt.Println("  install      Installs into currently attached EntandoApp the bundle in the current directory")
}

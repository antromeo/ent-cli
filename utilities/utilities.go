package utilities

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
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

func HttpGet(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request: ", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed with status code: ", response.Status)
		os.Exit(1)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response: ", err)
		os.Exit(1)
	}
	return data
}

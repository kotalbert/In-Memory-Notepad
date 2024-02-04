package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		run()
	}
}

func run() {
	fmt.Print("Enter a command and data:")
	userInput := getInput()
	command, data := parseInput(userInput)

	switch command {
	case "exit":
		fmt.Print("[Info] Bye!\n")
		os.Exit(0)
	default:
		fmt.Print(command, data)
	}
	fmt.Println(command, data)
}

func parseInput(input string) (string, string) {
	words := strings.Split(input, " ")
	return words[0], strings.Join(words[1:], " ")
}

// getInput gets the input from the user, which is a single line of text
func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

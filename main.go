package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// todo: add tests
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
	case "create":
		fmt.Print("Creating a new note with data: ", data)
	case "list":
		fmt.Print("Listing all notes")
	case "clear":
		fmt.Print("Clearing all notes")
	case "exit":
		exitProgram()
	default:
		fmt.Print(command, data)
	}
	fmt.Println(command, data)
}

func exitProgram() {
	fmt.Print("[Info] Bye!\n")
	os.Exit(0)
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

package main

import (
	"bufio"
	"fmt"
	"kotalbert/in-memory-notepad/notepad"
	"os"
	"strings"
)

func main() {
	npd := notepad.NewNotepad()
	for {
		fmt.Print("Enter a command and data:")
		userInput := getInput()
		command, data := parseInput(userInput)

		switch command {
		case "create":
			err := npd.CreateNote(data)
			if err != nil {
				fmt.Print("[Error] Notepad is full\n")
				continue
			}
			fmt.Print("[OK] The note was successfully created\n")
		case "list":
			npd.ListNotes()
		case "clear":
			npd.ClearNotes()
			fmt.Print("[OK] All notes were successfully deleted\n")
		case "exit":
			fmt.Print("[Info] Bye!\n")
			os.Exit(0)
		default:
			fmt.Print(command, data)
		}
	}
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

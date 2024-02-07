package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	npd := NewNotepad()
	for {
		fmt.Print("Enter a command and data: ")
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
			fmt.Print("[Error] Unknown command\n")
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

type Note struct {
	Id   int
	Text string
}

func (n Note) ToString() string {
	return fmt.Sprintf("[Info] %d: %s", n.Id, n.Text)
}

const maxNoteNumber = 5

type Notepad struct {
	notes []Note
}

func NewNotepad() *Notepad {
	return &Notepad{}
}

func (n *Notepad) CreateNote(text string) error {
	if len(n.notes) >= maxNoteNumber {
		return errors.New("the notepad is full")
	}
	n.notes = append(n.notes, Note{Id: len(n.notes) + 1, Text: text})
	return nil
}

// ToString returns a string representation of the notepad
//
//	That is, a numbered list of all notes
func (n *Notepad) ToString() string {
	var sb strings.Builder

	for _, note := range n.notes {
		sb.WriteString(note.ToString())
		sb.WriteString("\n")
	}
	return sb.String()
}

// ListNotes prints all notes to the console
func (n *Notepad) ListNotes() {
	fmt.Print(n.ToString())
}

func (n *Notepad) ClearNotes() {
	n.notes = []Note{}
}

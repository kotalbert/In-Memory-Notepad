package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter the maximum number of notes:")
	maxNoteNumber, err := strconv.Atoi(getInput())
	if err != nil {
		fmt.Println("[Error] Please enter a valid number")
		os.Exit(1)
	}
	npd := NewNotepad(maxNoteNumber)

	for {
		fmt.Print("Enter a command and data: ")
		userInput := getInput()
		command, data := parseInput(userInput)

		switch command {
		case "create":
			err := npd.CreateNote(data)
			if err != nil {
				fmt.Print(err.Error())
				continue
			}
			fmt.Print("[OK] The note was successfully created\n")
		case "update":
			noteId, newText := parseInput(data)
			noteIdInt, err := strconv.Atoi(noteId)
			if err != nil {
				fmt.Println("[Error] Please enter a valid number")
				continue
			}
			err = npd.UpdateNote(noteIdInt, newText)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "list":
			err := npd.ListNotes()
			if err != nil {
				fmt.Println(err.Error())
			}
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

type Notepad struct {
	notes         []*Note
	maxSize       int
	nextNoteIndex int
}

// NewNotepad creates a new notepad, after asking for maximum number of notes
func NewNotepad(maxNoteNumber int) *Notepad {
	notes := make([]*Note, maxNoteNumber)
	return &Notepad{maxSize: maxNoteNumber, notes: notes}
}

func (n *Notepad) CreateNote(text string) error {
	if n.nextNoteIndex >= n.maxSize {
		return errors.New("[Error] Notepad is full\n")
	}
	if len(text) == 0 {
		return errors.New("[Error] Missing note argument\n")
	}
	n.notes[n.nextNoteIndex] = &Note{Id: n.nextNoteIndex + 1, Text: text}
	return nil
}

// ToString returns a string representation of the notepad
//
//	That is, a numbered list of all notes
func (n *Notepad) ToString() string {
	var sb strings.Builder

	for _, note := range n.notes {
		if note == nil {
			continue
		}
		sb.WriteString(note.ToString())
		sb.WriteString("\n")
	}
	return sb.String()
}

// ListNotes prints all notes to the console
func (n *Notepad) ListNotes() error {
	if len(n.notes) == 0 {
		return errors.New("[Info] Notepad is empty\n")
	}
	fmt.Print(n.ToString())
	return nil
}

func (n *Notepad) ClearNotes() {
	n.notes = make([]*Note, 0, n.maxSize)
}

func (n *Notepad) UpdateNote(noteId int, newText string) error {
	if noteId < 1 || noteId > n.maxSize {
		msg := fmt.Sprintf("[Error] Position %d is out of the boundaries [1, %d]\n", noteId, n.maxSize)
		return errors.New(msg)
	}
	if n.notes[noteId-1] == nil {
		return errors.New("[Error] There is nothing to update\n")
	}
	n.notes[noteId-1] = &Note{Id: noteId, Text: newText}
	fmt.Printf("[OK] The note at position %d was successfully updated\n", noteId)
	return nil
}

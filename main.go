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
			if noteId == "" {
				fmt.Print("[Error] Missing position argument\n")
				continue
			}
			if newText == "" {
				fmt.Print("[Error] Missing note argument\n")
				continue
			}
			noteIdInt, err := strconv.Atoi(noteId)
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", noteId)
				continue
			}
			err = npd.UpdateNote(noteIdInt, newText)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("[OK] The note at position %d was successfully updated\n", noteIdInt)

		case "delete":
			noteId, _ := parseInput(data)
			if noteId == "" {
				fmt.Print("[Error] Missing position argument\n")
				continue
			}
			noteIdInt, err := strconv.Atoi(noteId)
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", noteId)
				continue
			}
			err = npd.DeleteNote(noteIdInt)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("[OK] The note at position %d was successfully deleted\n", noteIdInt)

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
	notes      []*Note
	maxSize    int
	noteNumber int
}

// NewNotepad creates a new notepad, after asking for maximum number of notes
func NewNotepad(maxNoteNumber int) *Notepad {
	notes := make([]*Note, 0, maxNoteNumber)
	return &Notepad{maxSize: maxNoteNumber, notes: notes}
}

// CreateNote creates a new note with the given text
//   - If the notepad is full, it returns an error
//   - If the text is empty, it returns an error
//   - Otherwise, it creates a new note and returns nil
//
// Keeps track of number of notes already created
func (n *Notepad) CreateNote(text string) error {
	if n.noteNumber >= n.maxSize {
		return errors.New("[Error] Notepad is full\n")
	}
	if len(text) == 0 {
		return errors.New("[Error] Missing note argument\n")
	}
	//n.notes[n.nextNoteIndex] = &Note{Id: n.nextNoteIndex + 1, Text: text}
	//n.nextNoteIndex++
	n.notes = append(n.notes, &Note{Id: len(n.notes) + 1, Text: text})
	n.noteNumber++
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
	if n.noteNumber == 0 {
		return errors.New("[Info] Notepad is empty\n")
	}
	fmt.Print(n.ToString())
	return nil
}

func (n *Notepad) ClearNotes() {
	n.notes = make([]*Note, 0, n.maxSize)
	n.noteNumber = 0
}

func (n *Notepad) UpdateNote(noteId int, newText string) error {
	if noteId < 1 || noteId > n.maxSize {
		msg := fmt.Sprintf("[Error] Position %d is out of the boundaries [1, %d]\n", noteId, n.maxSize)
		return errors.New(msg)
	}
	if len(n.notes) == 0 || n.notes[noteId-1] == nil {
		return errors.New("[Error] There is nothing to update\n")
	}
	if len(newText) == 0 {
		return errors.New("[Error] Missing note argument\n")
	}
	n.notes[noteId-1] = &Note{Id: noteId, Text: newText}
	return nil
}

func (n *Notepad) DeleteNote(noteId int) error {
	if noteId < 1 || noteId > n.maxSize {
		msg := fmt.Sprintf("[Error] Position %d is out of the boundaries [1, %d]\n", noteId, n.maxSize)
		return errors.New(msg)
	}
	if len(n.notes) == 0 || n.notes[noteId-1] == nil {
		return errors.New("[Error] There is nothing to delete\n")
	}
	n.notes[noteId-1] = nil
	n.notes = append(n.notes[:noteId-1], n.notes[noteId:]...)
	n.noteNumber--
	// update the note id's after the deleted note
	for i := noteId - 1; i < len(n.notes); i++ {
		n.notes[i].Id = i + 1
	}
	return nil
}

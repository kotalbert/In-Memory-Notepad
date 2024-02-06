package notepad

import (
	"errors"
	"fmt"
	"strings"
)

const maxNoteNumber = 5

type Notepad struct {
	notes []Note
}

func NewNotepad() *Notepad {
	return &Notepad{}
}

func (n *Notepad) CreateNote(text string) error {
	if len(n.notes) > maxNoteNumber {
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

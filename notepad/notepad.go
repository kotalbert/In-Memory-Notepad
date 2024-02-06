package notepad

import (
	"fmt"
	"strings"
)

type Notepad struct {
	notes []Note
}

func NewNotepad() *Notepad {
	return &Notepad{}
}

func (n *Notepad) CreateNote(text string) {
	n.notes = append(n.notes, Note{Id: len(n.notes) + 1, Text: text})
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

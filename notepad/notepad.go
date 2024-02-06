package notepad

import (
	"strings"
)

type Notepad struct {
	notes []Note
}

func NewNotepad() *Notepad {
	return &Notepad{}
}

func (n *Notepad) createNote(text string) {
	n.notes = append(n.notes, Note{Id: len(n.notes), Text: text})
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

func (n *Notepad) clearNotes() {
	n.notes = []Note{}
}

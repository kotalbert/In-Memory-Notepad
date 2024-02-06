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

func (n *Notepad) createNote(text string) {
	n.notes = append(n.notes, Note{Id: len(n.notes), Text: text})
}

func (n *Notepad) listNotes() string {
	var sb strings.Builder

	for _, note := range n.notes {
		sb.WriteString(fmt.Sprintf("%d: %s\n", note.Id, note.Text))
	}
	return sb.String()

}

func (n *Notepad) clearNotes() {
	n.notes = []Note{}
}

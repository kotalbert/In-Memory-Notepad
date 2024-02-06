package notepad

import (
	"fmt"
	"strings"
)

type Notepad struct {
	notes []Note
}

type Note struct {
	id   int
	text string
}

func (n Note) ToString() string {
	return fmt.Sprintf("[Info] %d: %s", n.id, n.text)
}

func NewNotepad() *Notepad {
	return &Notepad{}
}

func (n *Notepad) createNote(text string) {
	n.notes = append(n.notes, Note{id: len(n.notes), text: text})
}

func (n *Notepad) listNotes() string {
	var sb strings.Builder

	for _, note := range n.notes {
		sb.WriteString(fmt.Sprintf("%d: %s\n", note.id, note.text))
	}
	return sb.String()

}

func (n *Notepad) clearNotes() {
	n.notes = []Note{}
}

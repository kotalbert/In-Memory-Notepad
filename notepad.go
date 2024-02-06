package main

type Notepad struct {
	notes []Note
}

type Note struct {
	id   int
	text string
}

func New() *Notepad {
	return &Notepad{}
}

func (n *Notepad) createNote(text string) {
	n.notes = append(n.notes, Note{id: len(n.notes), text: text})
}

func (n *Notepad) listNotes() []Note {
	return n.notes
}

func (n *Notepad) clearNotes() {
	n.notes = []Note{}
}

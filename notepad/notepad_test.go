package notepad

import (
	"testing"
)

func TestNotepadConstructor(t *testing.T) {
	notepad := NewNotepad()
	if notepad == nil {
		t.Error("Expected notepad to be created")
	}
}

func TestCreateNote(t *testing.T) {
	notepad := NewNotepad()
	notepad.createNote("test")
	if len(notepad.notes) != 1 {
		t.Error("Expected note to be created")
	}
}

func TestCreatedNoteShouldHaveValidData(t *testing.T) {
	notepad := NewNotepad()
	notepad.createNote("test")
	if notepad.notes[0].Text != "test" {
		t.Error("Expected note to have valid data")
	}
}

func TestListNotes(t *testing.T) {
	notepad := NewNotepad()
	notepad.createNote("test1")
	notepad.createNote("test2")

}

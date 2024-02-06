package notepad

import (
	"fmt"
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
	notepad.CreateNote("test")
	if len(notepad.notes) != 1 {
		t.Error("Expected note to be created")
	}
}

func TestCreatedNoteShouldHaveValidData(t *testing.T) {
	notepad := NewNotepad()
	notepad.CreateNote("test")
	if notepad.notes[0].Text != "test" {
		t.Error("Expected note to have valid data")
	}
}

func TestNotepad_ToString(t *testing.T) {
	notepad := NewNotepad()
	notepad.CreateNote("test1")
	notepad.CreateNote("test2")

	expected := fmt.Sprintf("[Info] 1: test1\n[Info] 2: test2\n")

	if notepad.ToString() != expected {
		t.Error("Expected ToString to return a valid string")
	}

}

package main

import (
	"fmt"
	"testing"
)

func TestNote_ToString(t *testing.T) {
	note := Note{Id: 1, Text: "test"}
	if note.ToString() != "[Info] 1: test" {
		t.Error("Note.ToString() did not return the expected value")
	}

}

func TestNotepadConstructor(t *testing.T) {
	notepad := NewNotepad(5)
	if notepad == nil {
		t.Error("Expected notepad to be created")
	}
}

func TestCreateNote(t *testing.T) {
	notepad := NewNotepad(5)
	_ = notepad.CreateNote("test")
	if len(notepad.notes) != 1 {
		t.Error("Expected note to be created")
	}
}

func TestCreatedNoteShouldHaveValidData(t *testing.T) {
	notepad := NewNotepad(5)
	_ = notepad.CreateNote("test")
	if notepad.notes[0].Text != "test" {
		t.Error("Expected note to have valid data")
	}
}

func TestCreateNoteShouldReturnErrorIfNoteIsEmpty(t *testing.T) {
	notepad := NewNotepad(5)
	err := notepad.CreateNote("")
	if err == nil {
		t.Error("Expected CreateNote to return an error")
	}
}

func TestCreateNoteShouldRespectMaxSize(t *testing.T) {
	notepad := NewNotepad(1)
	_ = notepad.CreateNote("test")
	err := notepad.CreateNote("test")
	if err == nil {
		t.Error("Expected CreateNote to return an error")
	}
}

func TestNotepad_ToString(t *testing.T) {
	notepad := NewNotepad(5)
	_ = notepad.CreateNote("test1")
	_ = notepad.CreateNote("test2")

	expected := fmt.Sprintf("[Info] 1: test1\n[Info] 2: test2\n")

	if notepad.ToString() != expected {
		t.Error("Expected ToString to return a valid string")
	}

}

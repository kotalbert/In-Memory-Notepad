package main

import "testing"

func TestNotepadConstructor(t *testing.T) {
	notepad := New()
	if notepad == nil {
		t.Error("Expected notepad to be created")
	}
}

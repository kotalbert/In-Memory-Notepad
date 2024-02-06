package notepad

import "testing"

func TestNote_ToString(t *testing.T) {
	note := Note{Id: 1, Text: "test"}
	if note.ToString() != "[Info] 1: test" {
		t.Error("Note.ToString() did not return the expected value")
	}

}

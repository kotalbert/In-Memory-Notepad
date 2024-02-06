package notepad

import "fmt"

type Note struct {
	Id   int
	Text string
}

func (n Note) ToString() string {
	return fmt.Sprintf("[Info] %d: %s", n.Id, n.Text)
}

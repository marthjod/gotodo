package todotxt

import (
	"bufio"
	"io"
	"strings"

	"github.com/marthjod/gotodo/model/entry"
)

// TodoTxt represents a list of todo.txt file entries
type TodoTxt struct {
	Entries []entry.Entry
}

// Read ...
func Read(r io.Reader) *TodoTxt {
	var t = &TodoTxt{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		t.Entries = append(t.Entries, entry.Read(line))
	}

	return t
}

// String writes out TodoTxt in todo.txt format
func (t *TodoTxt) String() string {
	var concat = []string{}
	for _, e := range t.Entries {
		concat = append(concat, e.String())
	}

	return strings.Join(concat, "\n")
}

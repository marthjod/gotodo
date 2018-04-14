package todotxt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/marthjod/gotodo/model/entry"
)

// TodoTxt represents a list of todo.txt file entries
type TodoTxt struct {
	Entries []entry.Entry
}

// Read reads from Reader and maps to TodoTxt struct.
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

// Write writes TodoTxt to file.
func (t *TodoTxt) Write(path string) error {
	o, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer o.Close()

	return t.write(o)
}

func (t *TodoTxt) write(w io.Writer) error {
	_, err := io.WriteString(w, fmt.Sprintln(t))
	return err
}

// String renders TodoTxt in todo.txt format
func (t *TodoTxt) String() string {
	var concat = []string{}
	for _, e := range t.Entries {
		concat = append(concat, e.String())
	}

	return strings.Join(concat, "\n")
}

// JSON renders TodoTxt in JSON format
func (t *TodoTxt) JSON() []byte {
	js, err := json.MarshalIndent(t.Entries, "", "  ")
	if err != nil {
		return []byte(`{"error": "` + err.Error() + `"}`)
	}

	return js
}

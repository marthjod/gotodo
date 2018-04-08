package read

import (
	"bufio"
	"log"
	"os"

	"github.com/marthjod/gotodo/model/entry"

	"github.com/marthjod/gotodo/model/todotxt"
)

// Read reads in todo.txt file and converts to Go struct
func Read(todoFile string) (todotxt.TodoTxt, error) {
	var (
		err     error
		todo    todotxt.TodoTxt
		f       *os.File
		line    string
		scanner *bufio.Scanner
	)

	if _, err = os.Stat(todoFile); err != nil {
		log.Fatal(err)
	}

	if f, err = os.Open(todoFile); err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner = bufio.NewScanner(f)
	todo = todotxt.TodoTxt{}

	for scanner.Scan() {
		line = scanner.Text()

		if line == "" {
			continue
		}

		entry := entry.Read(line)
		todo.Entries = append(todo.Entries, entry)
	}

	if err = scanner.Err(); err != nil {
		return todotxt.TodoTxt{}, err
	}

	return todo, err
}

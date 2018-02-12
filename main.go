package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/marthjod/gotodo/model/entry"
	"github.com/marthjod/gotodo/model/todotxt"
)

// read in todo.txt file and convert to Go struct
func convert(todoFile string) (todotxt.TodoTxt, error) {
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

func main() {
	var (
		todo         todotxt.TodoTxt
		js           []byte
		todoFilename string
		err          error
	)

	flag.StringVar(&todoFilename, "t", "todo.txt", "todo.txt file to use")
	flag.Parse()

	if todo, err = convert(todoFilename); err != nil {
		log.Fatal(err)
	}

	js, err = json.MarshalIndent(todo.Entries, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", js)

}

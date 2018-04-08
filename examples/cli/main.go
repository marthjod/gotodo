package main

import (
	"flag"
	"log"

	"github.com/marthjod/gotodo/cli"
	"github.com/marthjod/gotodo/model/todotxt"
	"github.com/marthjod/gotodo/read"
)

func main() {
	var (
		todo         todotxt.TodoTxt
		err          error
		todoFilename = flag.String("t", "todo.txt", "todo.txt file to use")
	)

	flag.Parse()

	if todo, err = read.Read(*todoFilename); err != nil {
		log.Fatal(err)
	}

	cli.PrintPrefixed(&todo)
}

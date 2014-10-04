package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TodoTxtEntry struct {
	Contexts    []string
	Projects    []string
	Priority    string
	Description string
}

type TodoTxt struct {
	Entries []TodoTxtEntry
}

// write out in todo.txt format
func (t *TodoTxt) render() string {
	var (
		i, k     int
		entry    TodoTxtEntry
		entryStr string
		out      []string
	)

	out = make([]string, len(t.Entries))
	for i = 0; i < len(t.Entries); i++ {
		entry = t.Entries[i]
		entryStr = ""

		if entry.Priority != "" {
			entryStr += fmt.Sprintf("%v ", entry.Priority)
		}
		if len(entry.Projects) > 0 {
			for k = 0; k < len(entry.Projects); k++ {
				entryStr += fmt.Sprintf("%v ", entry.Projects[k])
			}
		}
		if len(entry.Contexts) > 0 {
			for k = 0; k < len(entry.Contexts); k++ {
				entryStr += fmt.Sprintf("%v ", entry.Contexts[k])
			}
		}

		entryStr += entry.Description
		out = append(out, entryStr)
	}

	return strings.Join(out, "\n")
}

// read in todo.txt file and convert to Go struct
func convert(todoFile string) (TodoTxt, error) {
	var (
		err        error
		todo       TodoTxt
		f          *os.File
		line       string
		scanner    *bufio.Scanner
		projects   []string
		contexts   []string
		priority   string
		projectsRE *regexp.Regexp
		contextsRE *regexp.Regexp
		priorityRE *regexp.Regexp
		i          int
	)

	projectsRE = regexp.MustCompile(`\+\w+`)
	contextsRE = regexp.MustCompile(`@\w+`)
	priorityRE = regexp.MustCompile(`\([A-Z]\)`)

	if _, err = os.Stat(todoFile); err != nil {
		log.Fatal(err)
	}

	if f, err = os.Open(todoFile); err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner = bufio.NewScanner(f)
	todo = TodoTxt{}

	for scanner.Scan() {
		line = scanner.Text()

		if line == "" {
			continue
		}

		contexts = contextsRE.FindAllString(line, -1)
		if contexts == nil {
			contexts = []string{}
		}
		projects = projectsRE.FindAllString(line, -1)
		if projects == nil {
			projects = []string{}
		}
		priority = priorityRE.FindString(line)

		// clear line of tokens already processed
		for i = 0; i < len(projects); i++ {
			line = projectsRE.ReplaceAllString(line, "")
		}
		for i = 0; i < len(contexts); i++ {
			line = contextsRE.ReplaceAllString(line, "")
		}
		line = priorityRE.ReplaceAllString(line, "")

		todo.Entries = append(todo.Entries, TodoTxtEntry{
			contexts,
			projects,
			priority,
			// remaining line contents = description
			strings.Trim(line, " "),
		})
	}

	if err = scanner.Err(); err != nil {
		return TodoTxt{}, err
	}

	//log.Printf("Marshaled %d entries\n", len(todo.Entries))

	return todo, err
}

func main() {
	var (
		todo         TodoTxt
		js           []byte
		err          error
		port         int
		todoFilename string
	)

	flag.IntVar(&port, "port", 4242, "Local port to listen on")
	flag.StringVar(&todoFilename, "todofile", "todo.txt", "todo.txt file to use")
	flag.Parse()

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {

		var (
			ret int
			err error
		)

		if todo, err = convert(todoFilename); err != nil {
			log.Fatal(err)
		}

		js, err = json.MarshalIndent(todo.Entries, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ret, err = w.Write(js)
		if err != nil {
			log.Printf("Return code %d, error %s\n", ret, err.Error())
		} else {
			log.Printf("Returned todo.txt JSON containing %d entries", len(todo.Entries))
		}
	})

	log.Printf("Serving on port %v\n", port)
	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

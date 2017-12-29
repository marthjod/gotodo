package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/marthjod/gotodo/entry"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// read in todo.txt file and convert to Go struct
func convert(todoFile string) (entry.TodoTxt, error) {
	var (
		err        error
		todo       entry.TodoTxt
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

	projectsRE = regexp.MustCompile(`\+[\w/]+`)
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
	todo = entry.TodoTxt{}

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

		todo.Entries = append(todo.Entries, entry.TodoTxtEntry{
			contexts,
			projects,
			priority,
			// remaining line contents = description
			strings.Trim(line, " "),
		})
	}

	if err = scanner.Err(); err != nil {
		return entry.TodoTxt{}, err
	}

	//log.Printf("Marshaled %d entries\n", len(todo.Entries))

	return todo, err
}

func main() {
	var (
		todo         entry.TodoTxt
		js           []byte
		err          error
		port         int
		todoFilename string
		contentDir   string
		path         string
		dir          os.FileInfo
	)

	flag.IntVar(&port, "port", 4242, "Local port to listen on")
	flag.StringVar(&todoFilename, "todofile", "todo.txt", "todo.txt file to use")
	flag.Parse()

	contentDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	if path, err = filepath.Abs(contentDir); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(3)
	}

	if dir, err = os.Stat(path); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(3)
	}

	if !dir.IsDir() {
		fmt.Printf("Not a directory: %v\n", dir.Name())
		os.Exit(3)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

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

	http.HandleFunc("/js/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	log.Printf("Serving on port %v\n", port)
	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

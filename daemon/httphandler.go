package daemon

import (
	"net/http"

	"github.com/marthjod/gotodo/cli"
	"github.com/marthjod/gotodo/model/todotxt"
)

// MethodHandler routes bases on HTTP method.
func MethodHandler(w http.ResponseWriter, r *http.Request, t *todotxt.TodoTxt) {
	switch r.Method {
	case "GET":
		GETHandler(w, r, t)
	}
}

// GETHandler routes GET requests based on requested format.
func GETHandler(w http.ResponseWriter, r *http.Request, t *todotxt.TodoTxt) {
	switch r.Header.Get("Accept") {
	case "application/json":
		JSONHandler(w, r, t)
	default:
		CLIHandler(w, r, t)
	}
}

// JSONHandler writes TodoTxt as JSON.
func JSONHandler(w http.ResponseWriter, r *http.Request, t *todotxt.TodoTxt) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(t.JSON())
}

// CLIHandler writes TodoTxt as string intended for CLI.
func CLIHandler(w http.ResponseWriter, r *http.Request, t *todotxt.TodoTxt) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(cli.Prefixed(t)))
}

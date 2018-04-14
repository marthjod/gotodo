package daemon

import (
	"net/http"

	"github.com/marthjod/gotodo/cli"
	"github.com/marthjod/gotodo/model/todotxt"
	"github.com/marthjod/gotodo/provider"
)

// FormatHandler routes GET requests based on requested format.
func FormatHandler(w http.ResponseWriter, r *http.Request, t *todotxt.TodoTxt) {
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

// UploadHandler uploads TodoTxt to provider/backend.
func UploadHandler(w http.ResponseWriter, r *http.Request, t *todotxt.TodoTxt, provider provider.Provider, remoteFile string, autoRename bool) {
	resp, err := provider.Upload(remoteFile, autoRename, t)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(resp))
}

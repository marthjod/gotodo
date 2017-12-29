package entry

import (
	"fmt"
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

package cli

import (
	"fmt"
	"sort"
	"strings"

	"github.com/marthjod/gotodo/model/entry"
	"github.com/marthjod/gotodo/model/todotxt"
)

// PrintPrefixed prints entries prefixed with line number.
func PrintPrefixed(t *todotxt.TodoTxt) {
	fmt.Print(Prefixed(t))
}

// Prefixed returns entries prefixed with line number and colored by priority.
func Prefixed(t *todotxt.TodoTxt) string {
	var result = []string{}
	sort.Sort(entry.ByPriority(t.Entries))
	for idx, entry := range t.Entries {
		result = append(result, fmt.Sprintf("%3d %s\n", idx, entry.ColorString()))
	}

	return strings.Join(result, "")
}

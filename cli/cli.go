package cli

import (
	"fmt"
	"sort"

	"github.com/marthjod/gotodo/model/entry"
	"github.com/marthjod/gotodo/model/todotxt"
)

// PrintPrefixed prints entries prefixed with line number.
func PrintPrefixed(t *todotxt.TodoTxt) {
	sort.Sort(entry.ByPriority(t.Entries))
	for idx, entry := range t.Entries {
		fmt.Printf("%3d  %s\n", idx, entry.ColorString())
	}
}

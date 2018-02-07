package context

import (
	"strings"
)

// Context represents a todo.txt context
type Context string

// GetContexts maps input string(s) to their corresponding Context(s)
func GetContexts(s ...string) []Context {
	var ret = []Context{}
	for _, t := range s {
		ret = append(ret, Context(strings.Trim(t, "@")))
	}
	return ret
}

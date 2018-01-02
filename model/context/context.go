package context

import (
	"strings"
)

// Context ...
type Context string

// GetContexts ...
func GetContexts(s ...string) []Context {
	var ret = []Context{}
	for _, t := range s {
		ret = append(ret, Context(strings.Trim(t, "@")))
	}
	return ret
}

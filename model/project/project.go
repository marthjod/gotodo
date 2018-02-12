package project

import (
	"strings"
)

// Project represents a todo.txt project
type Project string

// MatchRegexp defines the regular expression matching project strings
const MatchRegexp = `\+[\pL/]+`

// GetProjects maps input string(s) to their corresponding Project(s)
func GetProjects(s ...string) []Project {
	var ret = []Project{}
	for _, p := range s {
		ret = append(ret, Project(strings.Trim(p, "+")))
	}
	return ret
}

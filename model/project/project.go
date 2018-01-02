package project

import (
	"strings"
)

// Project ...
type Project string

// GetProjects ...
func GetProjects(s ...string) []Project {
	var ret = []Project{}
	for _, p := range s {
		ret = append(ret, Project(strings.Trim(p, "+")))
	}
	return ret
}

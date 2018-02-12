package project

import (
	"reflect"
	"testing"
)

var expected = []struct {
	projectStr []string
	projects   []Project
}{
	{
		projectStr: []string{},
		projects:   []Project{},
	},
	{
		projectStr: []string{"+__", "+--", "+@"},
		projects:   []Project{"__", "--", "@"},
	},
	{
		projectStr: []string{"+Foo"},
		projects:   []Project{"Foo"},
	},
	{
		projectStr: []string{"+Foo", "+Bar"},
		projects:   []Project{"Foo", "Bar"},
	},
	{
		projectStr: []string{"+Aldrei_for_ég_suður"},
		projects:   []Project{"Aldrei_for_ég_suður"},
	},
}

func TestGetProjects(t *testing.T) {
	for _, exp := range expected {
		actual := GetProjects(exp.projectStr...)
		if !reflect.DeepEqual(actual, exp.projects) {
			t.Errorf("Expected %+v, got %+v", exp.projects, actual)
		}
	}
}

package entry

import (
	"reflect"
	"testing"

	"github.com/marthjod/gotodo/model/context"
	"github.com/marthjod/gotodo/model/priority"
	"github.com/marthjod/gotodo/model/project"
)

var expected = []struct {
	entry    Entry
	rendered string
}{
	{
		entry: Entry{
			Contexts:    []context.Context{},
			Projects:    []project.Project{},
			Priority:    priority.None,
			Description: "minimal entry",
			Done:        false,
		},
		rendered: "minimal entry",
	},
	{
		entry: Entry{
			Contexts:    []context.Context{},
			Projects:    []project.Project{"foo"},
			Priority:    priority.None,
			Description: "I have a project",
			Done:        false,
		},
		rendered: "I have a project +foo",
	},
	{
		entry: Entry{
			Contexts:    []context.Context{"work"},
			Projects:    []project.Project{},
			Priority:    priority.None,
			Description: "I have a context",
			Done:        false,
		},
		rendered: "I have a context @work",
	},
	{
		entry: Entry{
			Contexts:    []context.Context{},
			Projects:    []project.Project{},
			Priority:    priority.A,
			Description: "I have a prio",
			Done:        false,
		},
		rendered: "(A) I have a prio",
	},
	{
		entry: Entry{
			Contexts:    []context.Context{"work"},
			Projects:    []project.Project{"foo"},
			Priority:    priority.B,
			Description: "I have a project, context, and prio",
			Done:        false,
		},
		rendered: "(B) I have a project, context, and prio +foo @work",
	},
	{
		entry: Entry{
			Contexts:    []context.Context{"work"},
			Projects:    []project.Project{"foo"},
			Priority:    priority.B,
			Description: "I am a finished task",
			Done:        true,
		},
		rendered: "x (B) I am a finished task +foo @work",
	},
}

func TestString(t *testing.T) {
	for _, exp := range expected {
		if exp.entry.String() != exp.rendered {
			t.Errorf("Expected %+v, got %+v", exp.rendered, exp.entry)
		}
	}
}

func TestRead(t *testing.T) {
	for _, exp := range expected {
		out := Read(exp.rendered)
		if !reflect.DeepEqual(out, exp.entry) {
			t.Errorf("Expected %+v, got %+v", exp.entry, out)
		}
	}
}

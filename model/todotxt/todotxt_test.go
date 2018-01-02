package todotxt

import (
	"os"
	"reflect"

	"github.com/marthjod/gotodo/model/context"
	"github.com/marthjod/gotodo/model/entry"
	"github.com/marthjod/gotodo/model/priority"
	"github.com/marthjod/gotodo/model/project"

	"testing"
)

var input = TodoTxt{
	Entries: []entry.Entry{
		{
			Contexts:    []context.Context{},
			Projects:    []project.Project{},
			Priority:    priority.None,
			Description: "minimal entry",
			Done:        false,
		},
		{
			Contexts:    []context.Context{},
			Projects:    []project.Project{"foo"},
			Priority:    priority.None,
			Description: "I have a project",
			Done:        false,
		},
		{
			Contexts:    []context.Context{"work"},
			Projects:    []project.Project{},
			Priority:    priority.None,
			Description: "I have a context",
			Done:        false,
		},
		{
			Contexts:    []context.Context{},
			Projects:    []project.Project{},
			Priority:    priority.A,
			Description: "I have a prio",
			Done:        false,
		},
		{
			Contexts:    []context.Context{"work"},
			Projects:    []project.Project{"foo"},
			Priority:    priority.B,
			Description: "I have a project, context, and prio",
			Done:        false,
		},
		{
			Contexts:    []context.Context{"work"},
			Projects:    []project.Project{"foo"},
			Priority:    priority.B,
			Description: "I am a finished task",
			Done:        true,
		},
	},
}

func TestString(t *testing.T) {
	expected := `minimal entry
I have a project +foo
I have a context @work
(A) I have a prio
(B) I have a project, context, and prio +foo @work
x (B) I am a finished task +foo @work`

	out := input.String()
	if out != expected {
		t.Errorf("%q does not match expected %q", out, expected)
	}
}

func TestRead(t *testing.T) {
	f, err := os.Open("testdata/valid.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer f.Close()

	out := Read(f)
	if !reflect.DeepEqual(out, &input) {
		t.Errorf("Expected %+v, git %+v", input, out)
	}

}

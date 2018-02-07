package context

import (
	"reflect"
	"testing"
)

var expected = []struct {
	contextStr []string
	contexts   []Context
}{
	{
		contextStr: []string{},
		contexts:   []Context{},
	},
	{
		contextStr: []string{"@__", "@--", "@+"},
		contexts:   []Context{"__", "--", "+"},
	},
	{
		contextStr: []string{"@Foo"},
		contexts:   []Context{"Foo"},
	},
	{
		contextStr: []string{"@Foo", "@Bar"},
		contexts:   []Context{"Foo", "Bar"},
	},
}

func TestGetContexts(t *testing.T) {
	for _, exp := range expected {
		actual := GetContexts(exp.contextStr...)
		if !reflect.DeepEqual(actual, exp.contexts) {
			t.Errorf("Expected %+v, got %+v", exp.contexts, actual)
		}
	}
}

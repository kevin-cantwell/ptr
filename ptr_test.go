package ptr_test

import (
	"testing"

	"github.com/kevin-cantwell/ptr"
)

type Test struct {
	Msg string
	Val any
}

func TestPtr(t *testing.T) {
	var i interface{}
	tests := []Test{
		{"string", "foo"},
		{"int64", 123},
		{"interface{}", i},
		{"*interface{}", &i},
		{"nil", nil},
	}
	for _, test := range tests {
		t.Run(test.Msg, func(t *testing.T) {
			p, v := &test.Val, test.Val
			if *ptr.P(v) != test.Val {
				t.Errorf("Failed to reference: %v", v)
			}
			if ptr.V(p) != v {
				t.Errorf("Failed to de-reference: %v", p)
			}
		})
	}
}

func TestPtrLiteral(t *testing.T) {
	var foo *string = ptr.P("bar")
	if *foo != "bar" {
		t.FailNow()
	}
}

package main

import (
	"reflect"
	"testing"
)

var simpletree, lefttree, righttree []*int

func init() {
	g := func(n int) *int {
		return &n
	}

	simpletree = []*int{g(50), g(60), g(70), nil, g(80), g(90), nil}

	lefttree = []*int{
		g(10),
		g(20), nil,
		g(30), nil, nil, nil,
		g(40), nil, nil, nil, nil, nil, nil, nil,
		g(50), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,

		g(60), g(70), // 50's children
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,

		nil, g(80), g(90), nil, // 60's and 70's children
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
	}

	righttree = []*int{
		g(10),
		nil, g(20),
		nil, nil, nil, g(30),
		nil, nil, nil, nil, nil, nil, nil, g(40),
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, g(50),

		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		g(60), g(70), // 50's children

		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, g(80), g(90), nil, // 60 and 70's children
	}
}

func TestMktreeAndSlice(t *testing.T) {
	f := func(t *testing.T, name string, input []*int) {
		t.Helper()

		t.Run(name, func(t *testing.T) {
			t.Helper()

			tree := Mktree(input)
			slice := tree.Slice()

			switch {

			// In all cases, slice(tree(slice)) == slice.
			// If the slice is nil or []*int{}, slice(tree(slice)) == nil.
			case len(input) == 0 && len(slice) == 0:
				return

			case !reflect.DeepEqual(slice, input):
				t.Errorf("Mktree returned unexpected tree\ntree = %v\nslice = %v", tree, slice)
			}
		})
	}

	f(t, "nil input", nil)
	f(t, "len=0 input", []*int{})
	f(t, "simple tree, some nodes absent", simpletree)
	f(t, "a simple tree at the end of a linked list (left)", lefttree)
	f(t, "a simple tree at the end of a linked list (right)", righttree)
}

func TestString(t *testing.T) {
	f := func(t *testing.T, name string, input *TreeNode, want string) {
		t.Helper()

		t.Run(name, func(t *testing.T) {
			t.Helper()

			if got := input.String(); got != want {
				t.Errorf("unexpected value: %v", got)
			}
		})
	}

	f(
		t,
		"no node == 0",
		nil,
		"()",
	)

	f(
		t,
		"no left == () (right . . . )",
		&TreeNode{
			Val:   10,
			Right: &TreeNode{Val: 20},
		},
		"(\n\t10\n\t() (20)\n)",
	)

	f(
		t,
		"no right == (left . . .) ()",
		&TreeNode{
			Val:  10,
			Left: &TreeNode{Val: 20},
		},
		"(\n\t10\n\t(20) ()\n)",
	)

	f(
		t,
		"no left no right == ([0-9]+)",
		&TreeNode{Val: 10},
		"(10)",
	)

	f(
		t,
		"complete node",
		&TreeNode{
			Val:   10,
			Left:  &TreeNode{Val: 20},
			Right: &TreeNode{Val: 30},
		},
		"(\n\t10\n\t(20)\n\t(30)\n)",
	)
}

func TestEqual(t *testing.T) {
	f := func(t *testing.T, name string, foo, bar []*int) {
		t.Helper()

		t.Run(name, func(t *testing.T) {
			t.Helper()

			if reflect.DeepEqual(foo, bar) != Equal(Mktree(foo), Mktree(bar)) {
				t.Error("trees must be equal/distinct, but Equal lies")
			}
		})
	}

	f(t, "equal trees must be equal", simpletree, simpletree)
	f(t, "distinct trees must be distinct", simpletree, lefttree)
}

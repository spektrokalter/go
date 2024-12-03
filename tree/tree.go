package main

import (
	"bufio"
	"fmt"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Mktree(arr []*int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(arr))[:0]
	for _, n := range arr {
		if n == nil {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, &TreeNode{Val: *n})
		}
	}

	// p is current row
	// p[i:] is next row
	// r is p[i:] iterator

	p := nodes
	for i := 1; len(p[i:]) != 0; i <<= 1 {
		r := p[i:]
		for range i {
			if p[0] != nil {
				p[0].Left = r[0]
				p[0].Right = r[1]
			}
			p = p[1:]
			r = r[2:]
		}
	}

	return nodes[0]
}

func (n *TreeNode) String() string {
	switch {
	case n == nil:
		return "()"
	case n.Left == nil && n.Right == nil:
		return fmt.Sprintf("(%d)", n.Val)
	}

	left := ind(n.Left.String())
	right := ind(n.Right.String())

	switch {
	case n.Left == nil:
		right = right[1:]
		return fmt.Sprintf("(\n\t%d\n\t() %s)", n.Val, right)
	case n.Right == nil:
		left = left[:len(left)-1]
		return fmt.Sprintf("(\n\t%d\n%s ()\n)", n.Val, left)
	default:
		return fmt.Sprintf("(\n\t%d\n%s%s)", n.Val, left, right)
	}
}

func (n *TreeNode) Slice() []*int {
	var out []*int

	q := []*TreeNode{n}
	qempty := func() bool {
		for _, n := range q {
			if n != nil {
				return false
			}
		}
		return true
	}

	for !qempty() {
		prevlen := len(q)
		for i := range prevlen {
			if q[i] == nil {
				out = append(out, nil)
				q = append(q, nil)
				q = append(q, nil)
				continue
			}

			out = append(out, &q[i].Val)
			q = append(q, q[i].Left)
			q = append(q, q[i].Right)
		}
		copy(q, q[prevlen:])
		q = q[:len(q[prevlen:])]
	}

	return out
}

func Equal(m, n *TreeNode) bool {
	return m == nil && n == nil || (m != nil && n != nil &&
		m.Val == n.Val &&
		Equal(m.Left, n.Left) && Equal(m.Right, n.Right))
}

func ind(s string) string {
	t := ""
	for scan := bufio.NewScanner(strings.NewReader(s)); scan.Scan(); t += "\t" + scan.Text() + "\n" {
	}
	return t
}

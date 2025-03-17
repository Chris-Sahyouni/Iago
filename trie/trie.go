package trie

import (
	"fmt"
	"iago/isa"
	"maps"
	"slices"
)

// Aho-Corasick Trie
type TrieNode struct {
	children    map[string]*TrieNode
	data        isa.Instruction
	failureLink *TrieNode
}

func Trie(inStream []isa.Instruction, isa isa.ISA) *TrieNode {
	root := buildTrie(inStream, isa)
	root.buildFailureLinks()
	return root
}

func reverse[T any](t []T) {
	l := 0
	r := len(t) - 1
	for l < r {
		t[l], t[r] = t[r], t[l]
		l++
		r--
	}
}

func newTrieNode(data isa.Instruction) *TrieNode {
	return &TrieNode{
		children:    map[string]*TrieNode{},
		data:        data,
		failureLink: nil,
	}
}

func buildTrie(inStream []isa.Instruction, isa isa.ISA) *TrieNode {
	reverse(inStream)

	fmt.Println("Should be a gadget terminator", inStream[0].Op)

	root := newTrieNode(inStream[0])

	curr := root

	for _, instr := range inStream {
		if instr.Op == isa.GadgetTerminator() {
			curr = root
			continue
		}
		child, ok := curr.children[instr.Op]
		if ok {
			curr = child
		} else {
			newChild := newTrieNode(instr)
			curr.children[instr.Op] = newChild
			curr = newChild
		}
	}
	return root
}

func (root *TrieNode) buildFailureLinks() {
	rootChildren := slices.Collect(maps.Values(root.children))
	root.failureLink = root
	for _, child := range rootChildren {
		child.failureLink = root
	}

	queue := rootChildren
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for op, child := range curr.children {
			fail := curr.failureLink
			for fail != root && !curr.hasChild(op) {
				fail = fail.failureLink
			}
			if fail.hasChild(op) {
				child.failureLink = fail.children[op]
			} else {
				child.failureLink = root
			}
			queue = append(queue, child)
		}
	}
}

func (t *TrieNode) hasChild(target string) bool {
	childOps := slices.Collect(maps.Keys(t.children))
	return slices.Contains(childOps, target)
}

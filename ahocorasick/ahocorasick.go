package ahocorasick

import (
	"fmt"
	"iago/exe"
	"iago/isa"
)

type trieNode struct {
	children map[string]*trieNode
	data exe.Instruction
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

func newTrieNode(data exe.Instruction) *trieNode {
	return &trieNode{
		children: map[string]*trieNode{},
		data: data,
	}
}


func newTrie(inStream []exe.Instruction, isa isa.ISA) *trieNode {
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


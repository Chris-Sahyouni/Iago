package trie

import (
	"iago/isa"
)

type Trie struct {
	children         []*Trie
	data             string
	instruction_size int // find a way to make this stored only once, its gonna blow up the memory footprint of the trie
}

func newTrie(isa isa.ISA) *Trie {
	return &Trie{
		children:         make([]*Trie, 0),
		data:             "",
		instruction_size: isa.InstructionSize(),
	}
}

func (t *Trie) insert(s []string) {
	curr_node := t
	curr_op := ""
	for i := 0; i < len(s); i++ {
		curr_op = s[i]
		for index, child := range curr_node.children {
			if child.data == curr_op {
				curr_node = curr_node.children[index]
				break
			}
		}
		new_node := &Trie{
			children:         make([]*Trie, 0),
			data:             curr_op,
			instruction_size: t.instruction_size,
		}
		curr_node.children = append(curr_node.children, new_node)
		curr_node = new_node
	}
}

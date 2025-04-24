package trie

import (
	"errors"
	"iago/src/isa"
	"maps"
	"slices"
	"fmt"
	"strconv"
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

func (t *TrieNode) Rop(target string, isa isa.ISA) ([]uint, error) {
	reverseTargetSequence, err := parseTarget(target, isa.InstructionSize())
	if err != nil {
		return nil, err
	}
	var gadgetAddrs []uint // built in reverse
	root := t
	curr := t

	var instr string
	i := 0
	for i < len(reverseTargetSequence) {
		instr = reverseTargetSequence[i]
		if curr.hasChild(instr) {
			curr = curr.children[instr]
			i++
		} else { // take failure link
			if curr == root {
				return nil, errors.New("insufficient gadgets to build target payload")
			}
			gadgetAddrs = append(gadgetAddrs, curr.data.Vaddr)
			curr = curr.failureLink
		}
	}
	gadgetAddrs = append(gadgetAddrs, curr.data.Vaddr)

	reverse(gadgetAddrs)
	return gadgetAddrs, nil
}

// All hex characters are 1 byte, so indexing directly into the string is Ok
func parseTarget(target string, instructionSize int) ([]string, error) {

	if len(target) % instructionSize != 0 {
		return nil, errors.New("malformed target: target length modulo instruction size not equal to 0")
	}

	hexCharsPerByte := 2
	var splitTarget []string
	for i := 0; i < len(target); i += (instructionSize * hexCharsPerByte) {
		splitTarget = append(splitTarget, target[i: i + (instructionSize * hexCharsPerByte)])
	}
	reverse(splitTarget)
	return splitTarget, nil
}

func (t *TrieNode) DrawTrie(addressRepBase int) {
	var lines []string
	currLevel := []*TrieNode{t}
	var newLevel []*TrieNode
	for len(currLevel) > 0 {
		var currLine string
		for _, n := range currLevel {
			currLine += "  " + n.data.Op + ":" + strconv.FormatUint(uint64(n.data.Vaddr), addressRepBase)
			newLevel = append(newLevel, slices.Collect(maps.Values(n.children))...)
		}
		for i := range len(lines) {
			lines[i] = "  " + lines[i]
		}
		lines = append(lines, currLine)
		currLevel = newLevel
		newLevel = make([]*TrieNode, 0)
	}
	for _, l := range lines {
		fmt.Println(l)
	}
}
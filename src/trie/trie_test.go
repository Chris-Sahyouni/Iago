package trie

import (
	"iago/src/isa"
	"testing"
)

// Ok making this recursive since it will only be used on the above Trie ever
func equals(t1 *TrieNode, t2 *TrieNode) bool {
	if t1.data != t2.data {
		return false
	}
	if len(t1.children) != len(t2.children) {
		return false
	}
	for key := range t1.children {
		_, ok := t2.children[key]
		if !ok {
			return false
		}
		if !equals(t1.children[key], t2.children[key]) {
			return false
		}
	}
	return true
}

func TestBuildTrie(t *testing.T) {

	var testInstructionStream = []isa.Instruction{
		{Op: "b", Vaddr: 1},
		{Op: "a", Vaddr: 2},
		{Op: "z", Vaddr: 3},

		{Op: "c", Vaddr: 4},
		{Op: "a", Vaddr: 5},
		{Op: "z", Vaddr: 6},

		{Op: "f", Vaddr: 7},
		{Op: "u", Vaddr: 8},
		{Op: "z", Vaddr: 9},
		{Op: "z", Vaddr: 10},
	}

	expectedTrie := &TrieNode{
		data:        testInstructionStream[9],
		failureLink: nil,
		children: map[string]*TrieNode{
			"u": &TrieNode{
				data:        testInstructionStream[7],
				failureLink: nil,
				children: map[string]*TrieNode{
					"f": &TrieNode{
						data:        testInstructionStream[6],
						failureLink: nil,
						children:    make(map[string]*TrieNode, 0),
					},
				},
			},
			"a": &TrieNode{
				data:        testInstructionStream[4],
				failureLink: nil,
				children: map[string]*TrieNode{
					"c": &TrieNode{
						data:        testInstructionStream[3],
						failureLink: nil,
						children:    make(map[string]*TrieNode, 0),
					},
					"b": &TrieNode{
						data:        testInstructionStream[0],
						failureLink: nil,
						children:    make(map[string]*TrieNode, 0),
					},
				},
			},
		},
	}

	actualTrie := buildTrie(testInstructionStream, isa.TestISA{})
	if !equals(actualTrie, expectedTrie) {
		t.Fail()
	}

}

func TestFailureLinks(t *testing.T) {

	var testInstructionStream = []isa.Instruction{
		{Op: "c", Vaddr: 1},
		{Op: "a", Vaddr: 2},
		{Op: "z", Vaddr: 3},

		{Op: "b", Vaddr: 4},
		{Op: "a", Vaddr: 5},
		{Op: "z", Vaddr: 6},

		{Op: "a", Vaddr: 7},
		{Op: "b", Vaddr: 8},
		{Op: "z", Vaddr: 9},
	}

	root := buildTrie(testInstructionStream, isa.TestISA{})
	root.buildFailureLinks()

	a := root.children["a"]
	b := root.children["b"]
	ab := a.children["b"]
	ba := b.children["a"]
	ac := a.children["c"]

	if ab.failureLink != b {
		t.Fail()
	}
	if ba.failureLink != a {
		t.Fail()
	}
	if ac.failureLink != root {
		t.Fail()
	}
	if root.failureLink != root {
		t.Fail()
	}

}

func TestRop(t *testing.T) {
	var testInstructionStream = []isa.Instruction{
		{Op: "i", Vaddr: 1},
		{Op: "a", Vaddr: 2},
		{Op: "g", Vaddr: 3},
		{Op: "o", Vaddr: 4},
		{Op: "z", Vaddr: 5},

		{Op: "o", Vaddr: 4},
		{Op: "t", Vaddr: 5},
		{Op: "h", Vaddr: 6},
		{Op: "e", Vaddr: 7},
		{Op: "l", Vaddr: 8},
		{Op: "l", Vaddr: 9},
		{Op: "o", Vaddr: 10},
		{Op: "z", Vaddr: 11},
	}

	root := buildTrie(testInstructionStream, isa.TestISA{})
	root.buildFailureLinks()

	var gAddrs []uint
	var err error

	gAddrs, err = root.Rop("iago", isa.TestISA{})
	if err != nil {
		t.Fail()
	}
	if len(gAddrs) != 1 || gAddrs[0] != 1 {
		t.Fail()
	}

	gAddrs, err = root.Rop("othello", isa.TestISA{})
	if err != nil {
		t.Fail()
	}
	if len(gAddrs) != 1 || gAddrs[0] != 10 {
		t.Fail()
	}

	gAddrs, err = root.Rop("iago", isa.TestISA{})
	if err != nil {
		t.Fail()
	}
	if len(gAddrs) != 1 || gAddrs[0] != 1 {
		t.Fail()
	}

	gAddrs, err = root.Rop("go", isa.TestISA{})
	if err != nil {
		t.Fail()
	}
	if len(gAddrs) != 1 || gAddrs[0] != 3 {
		t.Fail()
	}

	gAddrs, err = root.Rop("helloiago", isa.TestISA{})
	if err != nil {
		t.Fail()
	}
	if len(gAddrs) != 2 || gAddrs[0] != 6 || gAddrs[1] != 1 {
		t.Fail()
	}

	gAddrs, err = root.Rop("nothello", isa.TestISA{})
	if err == nil {
		t.Fail()
	}

}

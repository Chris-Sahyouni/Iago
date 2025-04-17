package trie

import (
	"fmt"
	"iago/src/isa"
	"maps"
	"reflect"
	"slices"
	"strconv"
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

func TestHasChild(t *testing.T) {
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

	root := buildTrie(testInstructionStream, isa.TestISA{})
	root.buildFailureLinks()

	if !root.hasChild("u") || !root.hasChild("a") {
		t.Fail()
	}
	if !root.children["a"].hasChild("b") || !root.children["a"].hasChild("c") {
		t.Fail()
	}
	if !root.children["u"].hasChild("f") {
		t.Fail()
	}
	if root.hasChild("q") {
		t.Fail()
	}
	if root.children["a"].children["b"].hasChild("q") {
		t.Fail()
	}

}

func TestParseTarget(t *testing.T) {
	var res []string
	var err error

	res, err = parseTarget("teststring", 1)
	if err != nil {
		t.Error("Error on instruction size 1 case")
	}
	if !reflect.DeepEqual(res, []string{"g", "n", "i", "r", "t", "s", "t", "s", "e", "t"}) {
		t.Error("Failed on instruction size 1 case")
	}

	res, err = parseTarget("teststring", 2)
	if err != nil {
		t.Error("Error on instruction size 2 case")
	}
	if !reflect.DeepEqual(res, []string{"ng", "ri", "st", "st", "te"}) {
		t.Error("Error on instruction size 2 case")
	}

	res, err = parseTarget("oddlength", 2)
	if err == nil {
		t.Error("No error when len(target) % instructionSize != 0")
	}
}

func TestRop(t *testing.T) {
	var testInstructionStream = []isa.Instruction{
		{Op: "i", Vaddr: 1},
		{Op: "a", Vaddr: 2},
		{Op: "g", Vaddr: 3},
		{Op: "o", Vaddr: 4},
		{Op: "z", Vaddr: 5},

		{Op: "o", Vaddr: 6},
		{Op: "t", Vaddr: 7},
		{Op: "h", Vaddr: 8},
		{Op: "e", Vaddr: 9},
		{Op: "l", Vaddr: 10},
		{Op: "l", Vaddr: 11},
		{Op: "o", Vaddr: 12},
		{Op: "z", Vaddr: 13},
	}

	root := buildTrie(testInstructionStream, isa.TestISA{})
	root.buildFailureLinks()

	root.drawTrie(10)

	var gAddrs []uint
	var err error

	gAddrs, err = root.Rop("iago", isa.TestISA{})
	if err != nil {
		t.Errorf("Error on target iago: %s", err)
	}
	if len(gAddrs) != 1 || gAddrs[0] != 1 {
		t.Errorf("Wrong gadgets on target: iago \n Expected: [1], Actual: %v\n", gAddrs)
	}

	gAddrs, err = root.Rop("othello", isa.TestISA{})
	if err != nil {
		t.Errorf("Error on target othello: %s", err)
	}
	if len(gAddrs) != 1 || gAddrs[0] != 6 {
		t.Errorf("Wrong gadgets on target: othello \n Expected: [6], Actual: %v\n", gAddrs)
	}

	gAddrs, err = root.Rop("go", isa.TestISA{})
	if err != nil {
		t.Errorf("Error on target go: %s", err)
	}
	if len(gAddrs) != 1 || gAddrs[0] != 3 {
		t.Errorf("Wrong gadgets on target: go \n Expected: [3], Actual: %v\n", gAddrs)
	}

	gAddrs, err = root.Rop("helloiago", isa.TestISA{})
	if err != nil {
		t.Errorf("Error on target helloiago: %s", err)
	}
	if len(gAddrs) != 2 || gAddrs[0] != 8 || gAddrs[1] != 1 {
		t.Errorf("Wrong gadgets on target: helloiago \n Expected: [8 1], Actual: %v\n", gAddrs)
	}

	gAddrs, err = root.Rop("nothello", isa.TestISA{})
	if err == nil {
		t.Error("No error on target: nothello")
	}

}

/* -------------------------------------------------------------------------- */
/*                                  Draw Trie                                 */
/* -------------------------------------------------------------------------- */

func (t *TrieNode) drawTrie(addressRepBase int) {
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

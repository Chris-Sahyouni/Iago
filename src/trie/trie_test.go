package trie

import (
	"iago/src/isa"
	"reflect"
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
		{Op: "b_", Vaddr: 1},
		{Op: "a_", Vaddr: 2},
		{Op: "z_", Vaddr: 3},

		{Op: "c_", Vaddr: 4},
		{Op: "a_", Vaddr: 5},
		{Op: "z_", Vaddr: 6},

		{Op: "f_", Vaddr: 7},
		{Op: "u_", Vaddr: 8},
		{Op: "z_", Vaddr: 9},
		{Op: "z_", Vaddr: 10},
	}

	expectedTrie := &TrieNode{
		data:        testInstructionStream[9],
		failureLink: nil,
		children: map[string]*TrieNode{
			"u_": &TrieNode{
				data:        testInstructionStream[7],
				failureLink: nil,
				children: map[string]*TrieNode{
					"f_": &TrieNode{
						data:        testInstructionStream[6],
						failureLink: nil,
						children:    make(map[string]*TrieNode, 0),
					},
				},
			},
			"a_": &TrieNode{
				data:        testInstructionStream[4],
				failureLink: nil,
				children: map[string]*TrieNode{
					"c_": &TrieNode{
						data:        testInstructionStream[3],
						failureLink: nil,
						children:    make(map[string]*TrieNode, 0),
					},
					"b_": &TrieNode{
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
		{Op: "c_", Vaddr: 1},
		{Op: "a_", Vaddr: 2},
		{Op: "z_", Vaddr: 3},

		{Op: "b_", Vaddr: 4},
		{Op: "a_", Vaddr: 5},
		{Op: "z_", Vaddr: 6},

		{Op: "a_", Vaddr: 7},
		{Op: "b_", Vaddr: 8},
		{Op: "z_", Vaddr: 9},
	}

	root := buildTrie(testInstructionStream, isa.TestISA{})
	root.buildFailureLinks()

	a := root.children["a_"]
	b := root.children["b_"]
	ab := a.children["b_"]
	ba := b.children["a_"]
	ac := a.children["c_"]

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
		{Op: "b_", Vaddr: 1},
		{Op: "a_", Vaddr: 2},
		{Op: "z_", Vaddr: 3},

		{Op: "c_", Vaddr: 4},
		{Op: "a_", Vaddr: 5},
		{Op: "z_", Vaddr: 6},

		{Op: "f_", Vaddr: 7},
		{Op: "u_", Vaddr: 8},
		{Op: "z_", Vaddr: 9},
		{Op: "z_", Vaddr: 10},
	}

	root := buildTrie(testInstructionStream, isa.TestISA{})
	root.buildFailureLinks()

	if !root.hasChild("u_") || !root.hasChild("a_") {
		t.Fail()
	}
	if !root.children["a_"].hasChild("b_") || !root.children["a_"].hasChild("c_") {
		t.Fail()
	}
	if !root.children["u_"].hasChild("f_") {
		t.Fail()
	}
	if root.hasChild("q_") {
		t.Fail()
	}
	if root.children["a_"].children["b_"].hasChild("q_") {
		t.Fail()
	}

}

func TestParseTarget(t *testing.T) {
	var res []string
	var err error

	res, err = parseTarget("teststringz_", isa.TestISA{})
	if err != nil {
		t.Error("Error on instruction size 1 case")
	}
	if !reflect.DeepEqual(res, []string{"ng", "ri", "st", "st", "te"}) {
		t.Error("Failed on instruction size 1 case")
	}

	res, err = parseTarget("aabbccddeeffzz__", isa.TestISA2{})
	if err != nil {
		t.Error("Error on instruction size 2 case")
	}
	if !reflect.DeepEqual(res, []string{"eeff", "ccdd", "aabb"}) {
		t.Error("Error on instruction size 2 case")
	}

	res, err = parseTarget("oddlengthzz__", isa.TestISA2{})
	if err == nil {
		t.Error("No error when len(target) % instructionSize != 0")
	}
}

func TestRop(t *testing.T) {
	var testInstructionStream = []isa.Instruction{
		{Op: "i_", Vaddr: 1},
		{Op: "a_", Vaddr: 2},
		{Op: "g_", Vaddr: 3},
		{Op: "o_", Vaddr: 4},
		{Op: "z_", Vaddr: 5},

		{Op: "o_", Vaddr: 6},
		{Op: "t_", Vaddr: 7},
		{Op: "h_", Vaddr: 8},
		{Op: "e_", Vaddr: 9},
		{Op: "l_", Vaddr: 10},
		{Op: "l_", Vaddr: 11},
		{Op: "o_", Vaddr: 12},
		{Op: "z_", Vaddr: 13},
	}

	root := buildTrie(testInstructionStream, isa.TestISA{})
	root.buildFailureLinks()

	root.DrawTrie(10)

	var gAddrs []uint
	var err error

	gAddrs, err = root.Rop("i_a_g_o_z_", isa.TestISA{})
	if err != nil {
		t.Errorf("Error on target i_a_g_o_z_: %s", err)
	}
	if len(gAddrs) != 1 || gAddrs[0] != 1 {
		t.Errorf("Wrong gadgets on target: i_a_g_o_z_ \n Expected: [1], Actual: %v\n", gAddrs)
	}

	gAddrs, err = root.Rop("o_t_h_e_l_l_o_z_", isa.TestISA{})
	if err != nil {
		t.Errorf("Error on target o_t_h_e_l_l_o_z_: %s", err)
	}
	if len(gAddrs) != 1 || gAddrs[0] != 6 {
		t.Errorf("Wrong gadgets on target: o_t_h_e_l_l_o_z_ \n Expected: [6], Actual: %v\n", gAddrs)
	}

	gAddrs, err = root.Rop("g_o_z_", isa.TestISA{})
	if err != nil {
		t.Errorf("Error on target g_o_z_: %s", err)
	}
	if len(gAddrs) != 1 || gAddrs[0] != 3 {
		t.Errorf("Wrong gadgets on target: g_o_z_ \n Expected: [3], Actual: %v\n", gAddrs)
	}

	gAddrs, err = root.Rop("h_e_l_l_o_i_a_g_o_z_", isa.TestISA{})
	if err != nil {
		t.Errorf("Error on target h_e_l_l_o_i_a_g_o_z_: %s", err)
	}
	if len(gAddrs) != 2 || gAddrs[0] != 8 || gAddrs[1] != 1 {
		t.Errorf("Wrong gadgets on target: h_e_l_l_o_i_a_g_o_z_ \n Expected: [8 1], Actual: %v\n", gAddrs)
	}

	gAddrs, err = root.Rop("n_o_t_h_e_l_l_o_z_", isa.TestISA{})
	if err == nil {
		t.Error("No error on target: n_o_t_h_e_l_l_o_z_")
	}

}

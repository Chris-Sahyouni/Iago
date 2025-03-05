package sfxtree

import (
	"iago/exe"
)

type SuffixTreeNode struct {
	children []*SuffixTreeNode
	incomingEdge []byte
}

func NewSuffixTree(instructionStream []exe.Instruction) *SuffixTreeNode {
	return nil
}
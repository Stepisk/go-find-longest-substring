package main

import (
	"fmt"
)

// SuffixTreeNode is a node in the suffix tree
type SuffixTreeNode struct {
	children map[rune]*SuffixTreeNode
}

// NewSuffixTreeNode creates a new suffix tree node
func NewSuffixTreeNode() *SuffixTreeNode {
	return &SuffixTreeNode{
		children: make(map[rune]*SuffixTreeNode),
	}
}

func (stn *SuffixTreeNode) Keys() []*SuffixTreeNode {
    keys := make([]*SuffixTreeNode, 0, len(stn.children))
    for _, key := range stn.children {
        keys = append(keys, key)
    }
    return keys
}

func BuildSuffixTree(text string) *SuffixTreeNode {
	tree := NewSuffixTreeNode()
	for i := 0; i < len(text); i++ {
		substring := text[i:]
		node := tree
		for _, letter := range substring {
			if _, ok := node.children[letter]; !ok {
				node.children[letter] = NewSuffixTreeNode()
			}
			node = node.children[letter]
		}
	}
	return tree
}

type Stack struct {
    data []interface{}
}

func NewStack() *Stack {
    return &Stack{
        data: make([]interface{}, 0),
    }
}

func (s *Stack) Len() int {
    return len(s.data)
}

func (s *Stack) Push(value interface{}) {
    s.data = append(s.data, value)
}

func (s *Stack) Pop() interface{} {
    back := s.data[len(s.data) - 1]
    s.data = s.data[:len(s.data) - 1]
    return back
}

// FindLongestRepeatingSubstring finds the longest repeating substring in a string using a suffix tree
func FindLongestRepeatingSubstring(text string) string {
	tree := BuildSuffixTree(text)
	maxLength := 0
	resultString := ""

	stack := make([]*SuffixTreeNode, 0, 10)
	stack = append(stack, tree)
    stackLen := len(stack)
	for stackLen != 0 {
        currNode := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
		keys := currNode.Keys()
		if len(keys) > 1 {
			maxLength = max(maxLength, len(keys))
			resultString = text[:maxLength]
		}
        stack = append(stack, keys...)
        stackLen = len(stack)
	}

	return resultString
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {

	//Create a string  with long repeating substring
	str := "abcabcab"

	// Call FindLongestRepeatingSubstring function to find the longest repeating substring
	longestSubstr := FindLongestRepeatingSubstring(str)

	// Print the result
	fmt.Println("The longest repeating substring is:", longestSubstr) // abc
}

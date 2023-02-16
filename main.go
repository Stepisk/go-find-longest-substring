package main

import (
	"fmt"
	"time"
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
    data []*SuffixTreeNode
}

func NewStack() *Stack {
    return &Stack{
        data: make([]*SuffixTreeNode, 0),
    }
}

func (s *Stack) Len() int {
    return len(s.data)
}

func (s *Stack) Push(value *SuffixTreeNode) {
    s.data = append(s.data, value)
}

func (s *Stack) PushSlice(values []*SuffixTreeNode) {
    s.data = append(s.data, values...)
}

func (s *Stack) Pop() *SuffixTreeNode {
    back := s.data[len(s.data) - 1]
    s.data = s.data[:len(s.data) - 1]
    return back
}

// FindLongestRepeatingSubstring finds the longest repeating substring in a string using a suffix tree
func FindLongestRepeatingSubstring(text string) string {
	tree := BuildSuffixTree(text)
	maxLength := 0
	resultString := ""

    stack := NewStack()
    stack.Push(tree)
	for stack.Len() != 0 {
        currNode := stack.Pop()
		keys := currNode.Keys()
		if len(keys) > 1 {
			maxLength = max(maxLength, len(keys))
			resultString = text[:maxLength]
		}
        stack.PushSlice(keys)
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
    timeStart := time.Now()
	longestSubstr := FindLongestRepeatingSubstring(str)
    elapsed := time.Since(timeStart)

	// Print the result
	fmt.Println("The longest repeating substring is:", longestSubstr) // abc
    fmt.Printf("Calculated in %s\n", elapsed)
}

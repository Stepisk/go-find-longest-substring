package main

import (
 "fmt"
 "strings"
)

// SuffixTreeNode is a node in the suffix tree
type SuffixTreeNode struct {
 children map[rune]*SuffixTreeNode
 suffix   string
}

// NewSuffixTreeNode creates a new suffix tree node
func NewSuffixTreeNode() *SuffixTreeNode {
 return &SuffixTreeNode{
  children: make(map[rune]*SuffixTreeNode),
 }
}

// AddSuffix adds a suffix to the suffix tree
func (stn *SuffixTreeNode) AddSuffix(suffix string) {
 if len(suffix) == 0 {
  return
 }

 // Get the first character of the suffix and see if it exists in the children of this node
 firstChar := rune(suffix[0])
 child, ok := stn.children[firstChar]

 // If not, create a new node for it and add it as a child of this node
 if !ok {
  child = NewSuffixTreeNode()
  stn.children[firstChar] = child
 }

 // Recursively add the rest of the suffix to the child node
 child.AddSuffix(suffix[1:])

 // Set the suffix of this node to the longest common suffix of all its children
 longestCommonSuffix := ""
 for _, child := range stn.children {
  if len(child.suffix) > len(longestCommonSuffix) {
   longestCommonSuffix = child.suffix
  }
 }

 stn.suffix = longestCommonSuffix
}

// FindLongestRepeatingSubstring finds the longest repeating substring in a string using a suffix tree
func FindLongestRepeatingSubstring(s string) string {
 // Create the root node of the suffix tree
 root := NewSuffixTreeNode()

 // Iterate over all suffixes of the string and add them to the suffix tree
 for i := 0; i < len(s); i++ {
  suffix := s[i:]
  root.AddSuffix(suffix)
 }

 // Find the longest common suffix of all the children of the root node and return it as the longest repeating substring
 longestRepeatingSubstring := ""
 for _, child := range root.children {
  if len(child.suffix) > len(longestRepeatingSubstring) {
   longestRepeatingSubstring = child.suffix
  }
 }

 return longestRepeatingSubstring
}

func main() {

    //Create a string  with long repeating substring 

    str := "ababababababab"

    // Call FindLongestRepeatingSubstring function to find the longest repeating substring 

    longestSubstr := FindLongestRepeatingSubstring(str)

    // Print the result 

    fmt.Println("The longest repeating substring is:", strings.Repeat(longestSubstr, 2))  // abab 
}

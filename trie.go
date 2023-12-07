package main

import (
	"fmt"
)

func createTrie() {
	fmt.Println()
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}
type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

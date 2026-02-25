package builtin

import (
	"os"
	"path/filepath"
)

type TrieNode struct {
	children  [256]*TrieNode
	endOfWord bool
}

func Trie() TrieNode {
	return TrieNode{}
}

func (T *TrieNode) insertWord(word string) {
	currentNode := T

	for _, char := range word {
		if currentNode.children[char] == nil {
			currentNode.children[char] = &TrieNode{}
		}
		currentNode = currentNode.children[char]
	}

	currentNode.endOfWord = true
}

func (T *TrieNode) findAllMatches(prefixString string) []string {
	currentNode := T

	// find the trieNode at the end of prefix string
	for _, char := range prefixString {
		if currentNode.children[char] == nil {
			return nil
		}
		currentNode = currentNode.children[char]
	}

	var matchedStrings []string

	// dfs the whole word set for matches
	var dfs func(T *TrieNode, currentString string)

	dfs = func(T *TrieNode, currentString string) {
		if T.endOfWord {
			matchedStrings = append(matchedStrings, prefixString+currentString)
		}
		for k, childNode := range T.children {
			if childNode == nil {
				continue
			}
			dfs(childNode, currentString+string(rune(k)))
		}
	}

	dfs(currentNode, "")

	return matchedStrings
}

var trie TrieNode = Trie()
var fileTrie TrieNode = Trie()

func buildTrie() {
	// Insert builtin commands
	for key := range CmdFuncMap {
		trie.insertWord(key)
	}

	// Insert files in PATH
	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			trie.insertWord(entry.Name())
		}
	}
}

func buildTrieForFiles() {
	// Inserts files in the current directory
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	entries, err := os.ReadDir(wd)
	if err != nil {
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		fileTrie.insertWord(entry.Name())
	}
}

func AutoComplete(prefixString string) []string {
	matchedStrings := trie.findAllMatches(prefixString)
	return matchedStrings
}

func FileAutoComplete(prefixString string) []string {
	matchedStrings := fileTrie.findAllMatches(prefixString)
	return matchedStrings
}

package builtin

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

func Trie() TrieNode {
	return TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

func (T *TrieNode) insertWord(word string) {
	currentNode := T

	for _, char := range word {
		if _, ok := currentNode.children[char]; !ok {
			currentNode.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		currentNode = currentNode.children[char]
	}

	currentNode.endOfWord = true
}

func (T *TrieNode) findAllMatches(prefixString string) []string {
	currentNode := T

	// find the trieNode at the end of prefix string
	for _, char := range prefixString {
		if _, ok := currentNode.children[char]; !ok {
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
			dfs(childNode, currentString+string(k))
		}
	}

	dfs(currentNode, "")

	return matchedStrings
}

var trie TrieNode = Trie()

func buildTrie() {
	for key := range CmdFuncMap {
		trie.insertWord(key)
	}
}

func AutoComplete(prefixString string) []string {
	matchedStrings := trie.findAllMatches(prefixString)
	return matchedStrings
}

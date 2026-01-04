package graph

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		index := ch - 'a'
		// 不存在就构建
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		// 然后移动下来
		node = node.children[index]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.SearchPrefix(prefix)
	return node != nil
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		index := ch - 'a'
		if node.children[index] == nil {
			return nil
		}
		node = node.children[index]
	}
	return node
}

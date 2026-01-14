package trie

type Node struct {
	children [26]*Node
	isEnd    bool
}
type Trie struct {
	root *Node
}

func Constructor() Trie {
	trie := Trie{
		root: &Node{},
	}
	return trie
}

func (this *Trie) Insert(word string) {
	curNode := this.root
	for _, ch := range word {
		index := ch - 'a'
		// 不存在就构建
		if curNode.children[index] == nil {
			curNode.children[index] = &Node{}
		}
		// 然后移动下来
		curNode = curNode.children[index]
	}
	curNode.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.SearchPrefix(prefix)
	return node != nil
}

func (this *Trie) SearchPrefix(prefix string) *Node {
	curNode := this.root
	for _, ch := range prefix {
		index := ch - 'a'
		if curNode.children[index] == nil {
			return nil
		}
		curNode = curNode.children[index]
	}
	return curNode
}

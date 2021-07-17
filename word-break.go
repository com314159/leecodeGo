type Node struct {
	IsWord bool
	NextNode map[rune]*Node
}

func (n *Node)findRuneNode(r rune) *Node{
	if n == nil {
		return nil
	}
	if v,ok := n.NextNode[r];ok {
		return v
	}
	return nil
}

func (n *Node)getNextNode(r rune) *Node {
	if n == nil {
		return nil
	}

	if v,ok := n.NextNode[r];ok {
		return v
	}
	newNode := Node{IsWord: false,NextNode: make(map[rune]*Node)}
	n.NextNode[r] = &newNode
	return &newNode
}

func initTrie(wordDict []string) *Node {
	var root *Node = &Node{IsWord: false,NextNode: make(map[rune]*Node)}
	p := root
	for _,s := range wordDict {
		rs := []rune(s)
		p = root
		for i,r := range rs {
			p = p.getNextNode(r)
			if i == len(rs)-1 {
				p.IsWord = true
			}
		}
	}
	return root
}
func checkWord(trie *Node,start int,rs []rune,record map[int]bool) bool{
	if v,ok := record[start]; ok {
		return v
	}
	if start>=len(rs) {
		record[start] = true
		return true
	}
	p := trie
	for i:=start;i<=len(rs)-1;i++ {
		node := p.findRuneNode(rs[i])
		if node == nil {
			record[start] = false
			return false
		}
		if node.IsWord && i==len(rs)-1{
			record[start] = true
			return true
		}
		if node.IsWord && checkWord(trie,i+1,rs,record) {
			record[start] = true
			return true
		}
		p=node
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	root := initTrie(wordDict)
	rs := []rune(s)
	record := make(map[int]bool,len(rs))
	return checkWord(root,0,rs,record)
}
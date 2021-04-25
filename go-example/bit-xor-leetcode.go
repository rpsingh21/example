package main

type Node struct {
	child [2]*Node
	cnt   int
}

func (node *Node) Insert(val, i int) {
	if i >= 0 {
		bit := (val >> i) & 1
		if node.child[bit] == nil {
			node.child[bit] = &Node{}
		}
		node.child[bit].cnt++
		node.child[bit].Insert(val, i-1)
	}
}

func (node *Node) Count(val, limit, i int) int {
	bit := (val >> i) & 1
	lbit := (limit >> i) & 1
	ans := 0
	if lbit == 1 {
		ans = node.child[bit].cnt
	}
	next := node.child[bit^lbit]
	if next != nil {
		return ans + next.Count(val, limit, i)
	}
	return ans
}

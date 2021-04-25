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
		if node.child[bit] == nil {
			return ans
		}
		ans = ans + node.child[bit].cnt
	}
	next := node.child[bit^lbit]
	if next != nil {
		return ans + next.Count(val, limit, i-1)
	}
	return ans
}

func main() {
	arr := [4]int{1, 4, 2, 7}
	low := 2
	high := 6
	root := Node{}
	ans := 0
	for _, v := range arr {
		ans = ans + root.Count(v, high+1, 15) - root.Count(v, low, 15)
		root.Insert(v, 15)
	}
	println(ans)
	t := '.'
	println(t, int('z'))
}

package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(data int) *Tree {
	if t.Root == nil {
		t.Root = &Node{data, nil, nil}
	} else {
		t.Root.Add_Node(data)
	}
	return t
}

func (n *Node) Add_Node(val int) {
	if n == nil {
		return
	}
	if val < n.value {
		if n.left == nil {
			n.left.value = val
			n.left.left = nil
			n.left.right = nil
		} else {
			n.Add_Node(val)
		}
	} else {
		if n.right == nil {
			n.right.value = val
			n.right.left = nil
			n.right.right = nil
		} else {
			n.Add_Node(val)
		}
	}

}

func main() {
	t := &Tree{}
	t.Insert(2)

}

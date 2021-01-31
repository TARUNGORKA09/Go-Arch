package main

type Node struct {
	val   int
	Left  *Node
	Right *Node
}
type Tree struct {
	Root *Node
}

func (t Tree) Insert_Tree(data int) {
	if t.Root == nil {
		t.Root = &Node{val: data, Left: nil, Right: nil}
	} else {
		t.Root.Insert_Node(data)
	}
}

func (n *Node) Insert_Node(data int) {
	if n == nil {
		return
	} else if data <= n.val {
		if n.Left == nil {
			n.Left = &Node{val: data, Left: nil, Right: nil}
		} else {
			n.Left.Insert_Node(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{val: data, Left: nil, Right: nil}
		} else {
			n.Right.Insert_Node(data)
		}
	}
}
func main() {

	tree := Tree{}
	tree.Insert_Tree(40)

}

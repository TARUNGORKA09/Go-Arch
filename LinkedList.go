package main

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	head *Node
	len  int
}

func (l LinkedList) Insert(data int) {
	n := Node{}
	if l.len == 0 {
		n = *l.head
		n.value = data
		n.next = nil
	} else {
		for n.next != nil {
			n.next = l.head
		}

	}

}
func (l LinkedList) Traverse() {

}

func main() {

}

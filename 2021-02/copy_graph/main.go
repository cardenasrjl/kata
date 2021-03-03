package main

import "fmt"

type Node struct {
	Data int
	Neighbors []*Node
}

func (n *Node) Add(node *Node) {
	n.Neighbors = append(n.Neighbors, node) 
}

func main() {
	n0 := &Node{Data: 0}
	n1 := &Node{Data: 1}
	n2 := &Node{Data: 2}
	n3 := &Node{Data: 3}
	n4 := &Node{Data: 4}
	
	n0.Add(n3)
	n0.Add(n4)
	
	n1.Add(n2)
	
	n3.Add(n2)

	n4.Add(n1)
	n4.Add(n3)
	n4.Add(n0)

	
	nodes := copy(n0, map[int]*Node{})
	fmt.Printf("%+v", nodes)
}


func copy(root *Node, dic map[int]*Node ) *Node {
	node, ok := dic[root.Data]
	if ok {
		return node
	}
	
	copyNode := &Node{Data: root.Data}
	dic[root.Data] = copyNode
	
	for _, n := range root.Neighbors {
		copyNode.Neighbors = append(copyNode.Neighbors,copy(n,dic))
	}
	
	return copyNode
} 
package main

import "fmt"

type Node struct {
	value int
	Right *Node
	Left *Node
}

type Stack struct {
	nodes []*Node
}

func (s *Stack) Push(n *Node) {
	s.nodes =append(s.nodes, n)
}

func (s *Stack) Pop() *Node {
	if len(s.nodes) == 0 {
		return nil
	}
	
	temp := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1] 
	
	return temp
}


type Queue struct {
	nodes []*Node
}

func (q *Queue) Push (n *Node) {
	if len(q.nodes) == 0 {
		q.nodes = append(q.nodes, n)
		return
	}
	
	tem := []*Node{n}
	q.nodes = append(tem, q.nodes...)
}

func (q *Queue) Pop () *Node {
	if len(q.nodes) == 0 {
		return nil
	}
	
	temp := q.nodes[len(q.nodes)-1]
	q.nodes = q.nodes[:len(q.nodes)-1]
	return temp
}



func main() {
	
	//s := []int{1,2,3,4,5}
	
	/*n1 := &Node{value: 1}
	n2 := &Node{value: 2}
	st := &Stack{}
	st.Push(n1)
	st.Push(n2)
	st.Pop()
	st.Pop()
	st.Pop()*/

	n1 := &Node{value: 1}
	n2 := &Node{value: 2}
	q := &Queue{}
	q.Push(n1)
	q.Push(n2)
	 q.Pop()
	removed := q.Pop()
	
	
	fmt.Printf("%+v", removed.value)
}


func copyGraph() {
	
}

func suma(n []int, target int) bool {
	nMap := map[int]bool{}
	for _,v := range n {
		_,exists := nMap[v]
		if exists {
			return true
		} else {
			nMap[target - v] = true
		}
	}
	return false
}


func missing(n []int) int {
	//fin the sum from n
	sumWithMissing := 0
	for i:=1;i<=len(n)+1;i++ {
		sumWithMissing+=i
	}
	
	//find the sum of the integers
	sumWithoutMissing := 0
	for _,v := range n {
		sumWithoutMissing+=v
	}
	return sumWithMissing -sumWithoutMissing
	
}

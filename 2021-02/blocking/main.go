package main

import "fmt"

/**
Implement locking in a binary tree. A binary tree node can be locked or unlocked only if all of its descendants or ancestors are not locked.

Design a binary tree node class with the following methods:

is_locked, which returns whether the node is locked
lock, which attempts to lock the node. If it cannot be locked, then it should return false. Otherwise, it should lock it and return true.
unlock, which unlocks the node. If it cannot be unlocked, then it should return false. Otherwise, it should unlock it and return true.
You may augment the node to add parent pointers or any other property you would like. You may assume the class is used in a single-threaded program, so there is no need for actual locks or mutexes. Each method should run in O(h), where h is the height of the tree.

 */


type Node struct {
	Value int
	Locked bool
	Parent *Node
	Left *Node
	Right *Node
}


func (n *Node) IsLocked() bool {
	return n.Locked
}

func (n *Node) AncestorsLocked() bool {
	if n != nil {
		if n.Locked {
			return true
		}
		return n.Parent.AncestorsLocked()
	}
	return false
}

func (n *Node) DescendantLocked() bool {
	if n != nil {
		if n.Locked {
			return true
		}
		return n.Left.DescendantLocked() || n.Right.DescendantLocked() 
	}
	
	return false
}


func (n *Node) Lock() bool {
	if !n.IsLocked() && !n.DescendantLocked() && !n.AncestorsLocked() { 
		n.Locked = true
		return true
	}
	return false
}

func (n *Node) Unlock() bool {
	if n.IsLocked() && !n.DescendantLocked() && !n.AncestorsLocked() {
		n.Locked =  false
		return true
	}
	return false
}


func main() {
	var parent *Node
	parent = &Node{
		Value:  1,
		Locked: false,
		Parent: nil,
		Left:   &Node{
			Value:  3,
		},
		Right:  &Node{
			Value:  0,
		},
	}
	parent.Left.Parent = parent
	parent.Right.Parent = parent
	
	locked := parent.Lock()
	locked2 := parent.Right.Lock()
	locked3 := parent.Left.Lock()
	fmt.Printf(" %t %t %t", locked, locked2, locked3)
}



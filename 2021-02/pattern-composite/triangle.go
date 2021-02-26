package main

import "fmt"

//Canvas ...
type Triangle struct {
	Name string
}

//Draw ...
func (t *Triangle) Draw() {
	fmt.Printf("Drawing a triangle %s \n", t.Name)
}
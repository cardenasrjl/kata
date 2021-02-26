package main

import "fmt"

//Canvas ...
type Circle struct {
	Name string
}

//Draw ...
func (t *Circle) Draw() {
	fmt.Printf("Drawing a circle %s \n", t.Name)
}
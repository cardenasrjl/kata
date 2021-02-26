package main

import "fmt"

//Canvas ...
type Canvas struct {
	Components []Component
	Name string
}

//Draw ...
func (c *Canvas) Draw() {
	fmt.Printf("Printing the canvas %s \n",c.Name )
	for _,f := range c.Components {
		f.Draw()
	}
}

func (c *Canvas) Add(component Component) {
	c.Components = append(c.Components, component)
}
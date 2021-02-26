package main

/**
From: https://golangbyexample.com/composite-design-pattern-golang/
We use this pattern when we want a group of objects called `composite` is treated in a similar way as a single object.

From Design pattersn: 
The intent is to compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly
 
When to use it:

* Makes sense to use in cases when the composite and individual objects needs to be treated in the same way from a client perspective.
* you want to represent part-whole hierarchies of objects
* you want clients to be able to ignore the difference between compositions of objects and individual objects. Clients will treat all objects in the composite structure uniformly.
*/


//Component ..
type Component interface {
	Draw()
}

func main() {
	can1 := Canvas{
		Name:       "This is the main canvas",
	}

	can2 := Canvas{
		Name:       "This is the sub canvas",
	}
	
	can1.Add(&can2)
	can1.Add(&Triangle{Name: "triangle2"})
	can2.Add(&Circle{Name: "Circle1"})
	can2.Add(&Circle{Name: "Circle2"})
	can2.Add(&Triangle{Name: "triangle1"})
	can1.Draw()
	
}








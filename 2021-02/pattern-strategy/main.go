package main

import "fmt"

/*
from golang by example:   
Strategy design pattern is a behavioral design pattern. This design pattern allows you to change the behavior of an object at run time without any change in the class of that object.
*/

//Data
type Data struct {
	s Sorting
	arr []int
}

//New ...
func  New(data []int, s Sorting) *Data {
	return &Data{
		s:   s,
		arr: data,
	}
}
//SetSortingAlgorithm ..
func (d *Data) SetSortingAlgorithm( s Sorting) {
	d.s = s
}

//SetSortingAlgorithm ..
func (d *Data) Sort() {
	d.s.SortInt(d.arr)
}

//SetSortingAlgorithm ..
func (d *Data) Print() {
	for _,v := range d.arr {
		fmt.Printf(" %d \n", v)
	}
}

//Sorting ...
type Sorting interface {
	SortInt(d []int ) 
}

//Bubble ...
type Bubble struct {
}

//Dijkstra ...
type Dijkstra struct {
}

//SortInt ...
func (b *Bubble) SortInt(d []int) {

	for i:=0;i<len(d)-1;i++ {
		for j := 0; j < len(d)-1; j++ {
			if d[i] > d[i+1] {
				temp := d[i]
				d[i] = d[i+1]
				d[i+1] = temp
			}

		}
	}
	
}

func (b *Dijkstra) SortInt(d []int) {
}

func main() {
	
	data := New([]int{100,1,2,3,4,5,6,7,8,9}, &Bubble{})
	data.Sort()
	data.Print()
	
	
}
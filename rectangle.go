package main

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) Area() int {
	return r.length * r.breadth
}

func NewRectangle(length int, breadth int) Rectangle {
	return Rectangle{length, breadth}
}

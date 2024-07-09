package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (x Rectangle) Area() int {
	return x.width * x.height
}

func main() {
	rect := Rectangle{
		width:  8,
		height: 10,
	}
	fmt.Println(rect.Area())
}

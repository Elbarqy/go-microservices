package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func (p *Point) decX() int32 {
	p.x--
	return p.x
}

func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	fmt.Println("---------------------------")
	fmt.Println("-----Working With Types----")
	p1 := Point{1, 2}
	fmt.Println(p1.x, p1.y)
	changeX(&p1)
	fmt.Println(p1.decX(), p1.y)
	fmt.Println("-----------END-------------")
}

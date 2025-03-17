package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

// Method tính khoảng cách từ gốc tọa độ (0,0) đến điểm
func (p Point) Distance() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// Method di chuyển điểm (ảnh hưởng đến giá trị gốc)
func (p *Point) Move(x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	p := Point{3, 4}

	fmt.Println(p.Distance())

	p.Move(2, -1)
	fmt.Println("Điểm sau khi di chuyển:", p) // {5,3}
}

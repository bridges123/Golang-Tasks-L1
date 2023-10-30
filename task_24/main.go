package main

import (
	"fmt"
	"math"
)

// Point описывает точку с координатами x,y типа float64
type Point struct {
	x float64
	y float64
}

// NewPoint возвращает ссылку на новый объект точки с переданными координатами
func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

// DistanceTo считает расстояние между текущей и переданной точкой
func (p *Point) DistanceTo(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p.x-p2.x, 2) + math.Pow(p.y-p2.y, 2))
}

func main() {
	p1, p2 := NewPoint(0, 0), NewPoint(3, 4)
	fmt.Printf("Distance: %f\n", p1.DistanceTo(p2))
}

package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(q, p Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - p.Y}
}

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(q.Distance(p))
	fmt.Println(p.Distance(q))

	// 三角形周长
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	//r := Point{1, 2}
	//r.ScaleBy(2)
	//fmt.Println(r)

	//r := &Point{1, 2}
	//r.ScaleBy(2)
	//fmt.Println(*r)

	r := Point{1, 2}
	(&r).ScaleBy(2)
	fmt.Println(r)

	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	var x = ColoredPoint{Point{1, 1}, red}
	var y = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(x.Distance(y.Point))
	x.ScaleBy(2)
	y.ScaleBy(2)
	fmt.Println(x.Distance(y.Point))

	var m = ColoredPoint2{&Point{1, 1}, red}
	var n = ColoredPoint2{&Point{5, 4}, blue}
	fmt.Println(m.Distance(*n.Point))
	m.ScaleBy(2)
	n.ScaleBy(2)
	fmt.Println(m.Distance(*n.Point))
	m.Point = n.Point
	fmt.Println(*m.Point, *m.Point)
}

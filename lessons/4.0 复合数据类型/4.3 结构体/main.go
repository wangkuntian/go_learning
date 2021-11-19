package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// 等价于返回&tree(value: value)
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t

}

func main() {
	type Point struct{ X, Y int }
	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}
	var w = Wheel{Circle{Point{8, 8}, 5}, 20}
	var w2 = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		}, Spokes: 20,
	}
	w.Y = 40
	w2.X = 42
	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n", w2)
}

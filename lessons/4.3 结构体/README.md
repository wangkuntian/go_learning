
# 结构体

```go
type Employee struct {
	ID           int
	Name Address string
	DoB          time.Time
	Position     string
	Salary       int
	ManagerID    int
}
var dilbert Employee
dilbert.Salary += 1000
position := &dilbert.Position
*position = "Senior" + *position
```
## 结构体字面量
结构体类型的值可以通过结构体字面量来设置。
```go
type Point struct {
	x, y int
}
p := Point{1, 2}
q := Point{x: 1, y: 2}
```
两种初始化方式不可以混合使用。
也无法使用第一种初始化方式来绕过不可导出变量无法在其他包中使用的规则。
```go
package p
type T struct {
	a, b int
}
```

```go
package q
import "p"
var _ = p.T{a: 1, b: 2}     // 编译错误，无法引用a、b
var _ = p.T{1, 2}           // 编译错误，无法引用a、b
```

```go
pp := &Point{1, 2}
```
等价于
```go
pp := new(Point)
*pp = Point{1, 2}
```
## 结构体嵌套和匿名成员

```go
type Circle struct {
	X, Y, Radius int
}
type Wheel struct {
	X, Y, Radius, Spokes int
}
```
转化成
```go
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}
```
可以将Circle和Wheel简化
```go
type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}
var w wheel
w.X = 8        // 等价于w.Circle.Point.X = 8
w.Y = 8        // 等价于w.Circle.Point.Y = 8
w.Radius = 5   // 等价于w.Circle.Radius = 5
w.Spokes = 20
```
结构体字面量必须遵循结构体的定义。
```go
w = Wheel{8, 8, 5, 20}                          // 编译错误，未知成员变量
w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20}    // 编译错误，未知成员变量

// 正确姿势
w = Wheel{Circle{Point{8, 8}, 5}, 20}
w = Wheel{Circle: Circle{Point: Point{X:8, Y: 8}, Radius: 5}, Spokes: 20}

```


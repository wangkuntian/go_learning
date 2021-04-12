
# 方法
```go
import (
    "fmt"
    "math"
)

type Point struct {
    X, Y float64
}

// 普通函数
func Distance(q, p Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 方法
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

p := Point{1, 2}
q := Point{4, 6}

// 调用
fmt.Println(Distance(p, q))
fmt.Println(q.Distance(p))
fmt.Println(p.Distance(q))

```
> 类型拥有的所有方法名都必须是唯一的，但不同的类型可以使用相同的方法名。

>  Go语言不允许为简单的内置类型添加方法
```go
// 方法非法，不能是内置数据类型
func (a int) Add (b int){    
    fmt.Println(a+b)
}
```

# 指针传递
由于主调函数会复制每一个实参变量，如果函数需要更新一个变量，或者实参变量太大而想避免复制整个实参，可以使用指针传递变量的地址。

```go
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

```
调用
```go
r := Point{1, 2}
r.ScaleBy(2)
fmt.Println(r)      // {2 4}
```
或者
```go
r := &Point{1, 2}
r.ScaleBy(2)
fmt.Println(*r)     // {2 4}
```
或者
```go
r := Point{1, 2}
(&r).ScaleBy(2)
fmt.Println(r)      // {2 4}
```

如果r是Point类型的变量（上面第一种），但是方法要求的是*Point类型，编译器会对变量进行隐式转换，从r转换成&r。

只有变量才允许这么做，包括结构体字段、数组或者slice元素。

不能对一个不能取地址的Point类型参数调用*Point方法，因为无法或者临时变量的地址。
```go
Point{1, 2}.ScaleBy(2) //编译错误。
```

# 结构体扩展
```go
type ColoredPoint struct {
	Point
	Color color.RGBA
}

red := color.RGBA{R: 255, A: 255}
blue := color.RGBA{B: 255, A: 255}
var x = ColoredPoint{Point{1, 1}, red}
var y = ColoredPoint{Point{5, 4}, blue}
fmt.Println(x.Distance(y.Point))        // 5
x.ScaleBy(2)
y.ScaleBy(2)
fmt.Println(x.Distance(y.Point))        // 10

```
匿名字段类型可以是指向命名类型的指针。
```go
type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

var m = ColoredPoint2{&Point{1, 1}, red}
var n = ColoredPoint2{&Point{5, 4}, blue}
fmt.Println(m.Distance(*n.Point))
m.ScaleBy(2)
n.ScaleBy(2)
fmt.Println(m.Distance(*n.Point))
m.Point = n.Point
fmt.Println(*m.Point, *m.Point)     // {10 8} {10 8}
```

# 方法变量
```go
p := Point(1, 2)
q := Point(4, 6)
distanceFromP := p.Distance         // 方法变量
fmt.Println(distanceFromP(q))       // 5

scaleP := p.ScaleBy                 // 方法变量
scaleP(2)                           // p(2, 4)
scaleP(3)                           // p(6, 12)
```

# 方法表达式
```go
p := Point(1, 2)
q := Point(4, 6)
distance := Point.Distance          // 方法表达式
fmt.Println(distance(p, q))         // 5

scale := (*Point).ScaleBy           // 方法表达式
scale(&p, 2)                        // p(2, 4)
scale(&p, 3)                        // p(6, 12)
```

# 封装
首字母大写的标识符是可以从包中导出的，相反的，小写的是不能被导出的。同样的机制也作用于结构体内的字段和类型方法。

> 要封装一个对象，必须使用结构体。

> 在Go语言中封装的单元是包而不是类型。无论在函数内的代码还是方法内的代码，结构体类型内的字段对于同一个包中的所有代码都是可见的。

封装提供了三个优点：
1. 使用方不能直接修改对象的变量，所以不需要更多的语句来检查变量的值。
2. 隐藏实现细节可以防止使用方依赖的属性发生改变，使设计更加灵活。
3. 防止使用者肆意改变对象内的变量。
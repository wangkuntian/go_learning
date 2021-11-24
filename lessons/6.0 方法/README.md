
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
r := &Point{1, 2}
r.ScaleBy(2)
fmt.Println(r)      // {2 4}
```
或者
```go
p := Point{1, 2}
pptr := &P
r.ScaleBy(2)
fmt.Println(*r)     // {2 4}
```
或者
```go
p := Point{1, 2}
(&p).ScaleBy(2)
fmt.Println(r)      // {2 4}
```

如果r是Point类型的变量（上面第一种），但是方法要求的是*Point类型，编译器会对变量进行隐式转换，从r转换成&r。

只有变量才允许这么做，包括结构体字段、数组或者slice元素。

不能对一个不能取地址的Point类型参数调用*Point方法，因为无法或者临时变量的地址。
```go
Point{1, 2}.ScaleBy(2) //编译错误。
```

在合法的方法调用表达式中，只有符合下面三种形式的语句才能够成立。
1. 实参接收者和形参接收者是同一个类型，比如都是T类型或者都是*T类型。
```go
Point{1, 2}.Distance(q) // Point
pptr.ScaleBy(2)         // *Point
```

2. 实参接收者是T类型的变量而形参接收这个是*T类型。编译器会隐式地获取变量的地址。
```go
p.ScaleBy(2)            // 隐式转换为(&p)
```

3. 实参接收者是*T类型而形参接收者是T类型。编译器会隐式地引用接收者，获得实际的值。
```go
pptr.Distance(q)        // 隐式转换为(*pptr)
```

如果所有类型T方法的接收者是T自己（而非*T），那么复制它的实例是安全的；调用方法的时候都必须进行一次复制。但是任何方法的接收者是指针的情况下，应该避免复制T的实例，因为这么做可能会破坏内部原本的数据。

```go
Point{1, 2}.ScaleBy(2)  // 编译错误，无法从一个无地址的Point值上调用ScaleBy。
```

## nil是一个合法的接收者
就像一些函数允许nil指针作为实参，方法接收者也一样，尤其是当nil是类型中有意义的零值（如map和slice类型）时，更是如此。

在这个简单的整型数链表中，nil代表空链表。
```go
type IntList struct {
	Value int
	Tail *IntList
}
// Sum返回列表元素的总和
func (list *IntList) Sum() int {
	if list == nil {
		return 0
    }
	return list.Value + list.Tail.Sum()
}
```

# 结构体内嵌
```go
import "image/color"

type Point struct{ X, Y float64 }

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
p := Point{1, 2}
q := Point{4, 6}
distanceFromP := p.Distance         // 方法变量
fmt.Println(distanceFromP(q))       // 5

scaleP := p.ScaleBy                 // 方法变量
scaleP(2)                           // p(2, 4)
scaleP(3)                           // p(6, 12)
```

# 方法表达式
```go
p := Point{1, 2}
q := Point{4, 6}
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
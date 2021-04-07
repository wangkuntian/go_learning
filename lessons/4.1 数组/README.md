
# 数组
数组是具有固定长度且拥有0个或多个相同数据类型元素的序列。Go里很少用到。
```go
var a [3]int                // [0, 0, 0]
fmt.Println(a[0])           // 0
fmt.Println(a[len(a) - 1])  // 0

var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
fmt.Println(r[2])           // 0

var l = [...]int{1, 2, 3}
fmt.Printf("%T \n", l)      // [3]int
```
数组初始化的值为元素类型的零值。

数组长度是数组类型的一部分，所以[3]int和[4]int是两种不同的数组类型。

数组的长度必须是常量表达式，表达式的值在程序编译是就可以确定。

```go
q := [3]int{1, 2, 3}
q = [4]int{1, 2, 3, 4}      // 编译错误
```

```go
r := [...]int{ 99: -1}      // 除了最后一个数是-1，其他为0。 
                            // 数组长度为100。

```

如果数组的元素类型是可比较的，那么这个数组也是可比较的。可以使用==或者!=比较两个数组。

# 切片slice
slice表示一个拥有相同类型的可变长度的序列。
slice通常写成[]T，其中T为元素俺的类型。

slice有三个属性：指针、长度和容量。
指针指向可以从slice中访问的第一个元素。
长度是指slice元素的个数，它不能超过slice的容量。
容量的大小通常指从slice的起始元素和底层数组的最后元素间的元素个数。

len函数获取slice的长度。
cap函数获取slice的容量。

slice无法作比较，可自己实现函数。
```go
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
    }
    for i := range x {
    	if x[i] != y[i] {
    		return fasle
        }   
    }
    return true
}
```

slice类型的零值是nil。
值为nil的slice没有对应的底层数组，它的长度和容量都是0。

```go
var s []int         // len(s) == 0, s == nil
s = nil             // len(s) == 0, s == nil
s = []int(nil)      // len(s) == 0, s == nil
s = []int{}         // len(s) == 0, s != nil
```
> 检查slice是否为空，使用len(s) == 0，而不是s == nil。

## make函数
make函数可以创建一个具有指定元素类型、长度和容量的slice。其中容量参数可以省略，这时，slice的长度和容量相等。
```go
make([]T, len)
make([]T, len, cap) // 和make([]T, cap)[:len]功能相同。
```
> make函数创建了一个无名数组并返回了它的一个slice。

## append函数
append函数用来将元素追加到slice的后面。

# 切片与数组的区别
1. 切片不是数组，但是切片底层指向数组
2. 切片本身长度是不一定的因此不可以比较，数组是可以的。
3. 切片是变长数组的替代方案，可以关联到指向的底层数组的局部或者全部。
4. 切片是引用传递（传递指针地址)，而数组是值传递（拷贝值）。
5. 切片可以直接创建，引用其他切片或数组创建。
6. 如果多个切片指向相同的底层数组，其中一个值的修改会影响所有的切片。
# 声明

```go
var name type = expression
```

类型和表达式部分可以省略一个，但是不能都省略。如果类型省略，它的类型将由初始化代表式决定；如果表达式省略，其初始值对应于类型的零值。
对于数字是0，对于布尔值是false，对于字符串是""，对于接口和引用类型（slice、指针、map、通道、函数）是nil。
对于一个像数组或结构体这样的复合类型，零值是所有元素或成员的零值。

## 简化
```go
var s sting
var i, j, k int
var b, f, s = true, 2.3, "four"
var f, error = os.Open(name)
```

# 短变量
> name := expression
> name的类型由expression的类型决定
> 短变量声明不需要声明所有在左边的变量。
```go
t := 0.0
i, j := 0, 1
i, j = j, i     // 交换i和j的值
f, err := os.Open(name)
```
短变量至少要声明一个新的变量，否则编译出错。
```go
f, err := os.Open(name)
f, err := os.Open(name)  // 出错
```

# 指针
```go
x := 1
p := &x         // p是整型指针( *int)，指向x
fmt.Println(*p) // "1"
*p = 2          // x = 2
fmt.Println(x)  // "2"
```

指针类型的零值是nil。
```go
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
```

```go
func incr(p * int) int {
    *p++
    return *p
}
v := 1
incr(&v)    // v == 2
incr(&v)    // v == 3
```

# new函数
```go
p := new(int)   // *int类型的p
*p              // "0"
*p = 2     
*p              // "2"
```
表达式new(T)创建一个未命名的T类型变量，初始化为T类型的零值，并返回其地址（地址类型为*T）。

# 变量生命周期
生命周期是指在程序执行过程中变量存在的时间段。变量的生命周期是通过它是否可达来确定的。
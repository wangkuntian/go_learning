# 声明
```go
var name type = expression
```

## 简化
```go
var s sting
var i, j, k int
var b, f, s = true, 2.3, "four"
var f, error = os.Open(name)
```

# 短变量
```go
t := 0.0
i, j := 0, 1
i, j = j, i     // 交换i和j的值
f, err := os.Open(name)
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

# 变量生命周期
变量的生命周期是通过它是否可达来确定的。
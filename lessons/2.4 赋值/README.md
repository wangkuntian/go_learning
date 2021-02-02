
```go
x = 1
*p = true
person.name = "John"
count[x] = count[x] * scale
count[x] *= scale
v := 1
v++
v--
```

# 多重赋值

```go
x, y = y, x
a[i], a[j] = a[j], a[i]
```

两个整数的最大公约数
```go
func gcd(x, y int) int {
    for y != 0 {
        x, y = x, x % y
    }
    return x
}
```

计算斐波那契数列的第N个数
```go
func fib(n int) int {
    x , y = 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x + y
    }
    return x
}
```

```go
f, err = os.Open(name)
_, err = os.Open(name)
```
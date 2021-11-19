
# 函数
# 声明
```go
func name(paramter-list) (result-list){
	body
}
```
例如：
```go
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
```
> 形参或者返回值得类型相同，那么类型只需要写一次。

函数的类型称作函数签名。当两个函数拥有相同的形参列表和返回列表时，认为这个两个函数的类型或签名是相同的。

> 实参是按值传递的。
> 实参包含引用类型，如指针、slice、map、函数或者通道，那么函数使用形参变量时可能会间接修改实参变量。

有些函数没有函数体，那么说明这个函数使用了除Go以外的语言实现。
```go
package math
func Sin(x float64) float64 // 使用汇编语言实现
```

# 多返回值
一个函数可以返回不只一个结果。

一个函数如果有命名的返回值，可以省略return语句的操作数，这称为裸返回。

```go

func countWordsAndImages(n *html.Node) (words, images int) {}

func CountWordsAndImages(url string) (words, images int, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(response.Body)
	response.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
```
裸返回是将每个命名返回结果按照顺序返回的快捷方式，所以在上面的函数中，每个return语句都等同于：
```go
return words, images, err
```

# 函数变量
函数可以赋给变量或者传递给其他函数或者作为返回值返回。

# 匿名函数

# 变长函数
在参数的之后参数类型之前使用"..."表示声明一个变长函数，调用这个函数时可以传递任意数目的参数。
```go
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
        total += val
	}
	return total
}
fmt.Println(sum())                // 0
fmt.Println(sum(3))               // 3
fmt.Println(sum(1, 2, 3, 4))      // 10
values := []int{1, 2, 3, 4}
fmt.Println(sum(value...))        // 10
```

# 延迟函数调用
defer语句是在return之前执行的。
```go
func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

fmt.Println(f())        // 1
fmt.Println(f2())       // 5
fmt.Println(f3())       // 1
```

# 宕机
Go语言的类型系统会捕获许多编译时错误，但有些错误（数组越界访问或引用空指针）需要在运行时进行检查。当Go语言运行时检测到这些错误，它就会发生宕机。
当宕机发生时，所有的延迟函数以倒序执行，从栈最上面的函数开始一直返回至main函数。
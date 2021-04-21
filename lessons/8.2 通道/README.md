# 通道
如果说goroutine是Go程序并发的执行体，通道就是它们之间的连接。通道是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

每一个通道是一个具体类型的管道，叫作通道的元素类型。一个int类型元素的通道写为chan int。

使用内置函数make来创建一个通道。
```go
ch := make(chat int)    // ch的类型是'chan int'
```

通道和map一样都是引用类型。通道类型的零值是nil。
同种类型的通道可以使用==比较。当两者都是同一通道数据的引用时，比较值为true。通道也可以和nil相比较。

通道有两种主要操作：发送（send）和接收（receive），二者统称为通信。
send语句从一个goroutine传输一个值到另一个执行接收表达式的goroutine。
两种操作都使用 <- 操作符。
```go
ch <- x     // 发送语句
x = <- ch   // 复制语句中的接收表达式
<-ch        // 接收语句，丢弃结果
```

通道支持第三个操作：关闭（close），它设置一个标志位来指示值当前已经发送完毕，这个通道后面已经没有值了。关闭后的发送操作将导致宕机。
在一个已经关闭的通道上进行接收操作，将获取所有已经发送的值，直到通道为空；这时任何接收操作会立即完成，同时获取到一个通道元素类型对应的零值。

调用内置的close函数关闭通道。
```go
close(ch)
```

使用简单的make调用创建的通道叫无缓冲（unbuffered）通道。make还可以接收第二个可选参数，一个表示通道的容量的整数。如果容量是0，make创建一个无缓冲通道。
```go
ch = make(chan int)         // 无缓冲通道
ch = make(chan int, 0)      // 无缓冲通道
ch = make(chan int, 3)      // 容量为3的缓冲通道
```

# 无缓冲通道
无缓冲通道的发送操作将会阻塞，直到另一个goroutine在对应的通道上执行接收操作，这时值传送完成，两个goroutine都可以继续执行。相反，如果先执行接收操作，接收方goroutine将阻塞，直到另一个goroutine在同一个通道上发送一个值。

使用无缓冲通道进行的通信会导致发送和接收goroutine同步化，因此，无缓冲通道也成为同步通道。

当一个值在无缓冲通道上传递时，接收值后发送方的goroutine才会再次被唤醒。

# 缓冲通道
创建一个可以容纳3个字符串的缓冲通道。
```go
ch = make(chan string, 3)
```

缓冲通道上的发送操作会在队列的尾部插入一个元素，接收操作从队列的头部移除一个元素。

如果通道满了，发送操作会阻塞所在的goroutine直到另一个goroutine对它进行操作来留可用的空间。如果通道是空的，执行接收操作的goroutine会阻塞，直到另一个goroutine在通道上发送数据。

## 容量
可以使用内置函数cap获取缓冲通道的容量。
```go
fmt.Println(cap(ch))
```
可以使用内置函数len获取当前通道内的元素个数，不过由于在并发程序中这个信息会随着检索操作很快过时，不常用。
```go
fmt.Println(len(ch))
```

## 示例

```go
func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses  // 返回最快的response。
}

func request(hostname string) (response string) {}
```
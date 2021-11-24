# 概述
Go有两种并发编程的风格。goroutine和通道（channel），它们支持通信顺序进程（Communication Sequential
Process，CSP），CSP是一个并发的模式，在不同的执行体（goroutine）之间传递值，但是变量本身局限于单一的执行体。

# goroutine
在Go里，每一个并发执行的活动称为goroutine。

当一个程序启动时，只有一个goroutine来调用main函数，称它为主routine。新的goroutine通过go语句进行创建。
```go
f()    // 调用f()；等待它返回
go f() // 新建一个调用f()的goroutine，不用等待。
```

# 通道
如果说goroutine是Go程序并发的执行体，通道就是它们之间的连接。通道是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
每一个通道是一个具体类型的管道，叫作通道的元素类型。一个int类型元素的通道写为chan int。

使用内置函数make来创建一个通道。
```go
ch := make(chat int) // ch的类型是'chan int'
```

## 通道的发送和接收

通道和map一样都是引用类型。通道类型的零值是nil。 同种类型的通道可以使用==比较。当两者都是同一通道数据的引用时，比较值为true。通道也可以和nil相比较。

通道有两种主要操作：发送（send）和接收（receive），二者统称为通信。 send语句从一个goroutine传输一个值到另一个执行接收表达式的goroutine。 两种操作都使用 <-
操作符。
```go
ch <- x   // 发送语句
x = <- ch // 赋值语句中的接收表达式
<-ch // 接收语句，丢弃结果
```

## 通道的关闭

通道支持第三个操作：关闭（close），它设置一个标志位来指示值当前已经发送完毕，这个通道后面已经没有值了。关闭后的发送操作将导致宕机。
在一个已经关闭的通道上进行接收操作，将获取所有已经发送的值，直到通道为空；这时任何接收操作会立即完成，同时获取到一个通道元素类型对应的零值。

调用内置的close函数关闭通道。
```go
close(ch)
```
## 通道的分类
### 无缓冲通道
使用简单的make调用创建的通道叫无缓冲（unbuffered）通道。make还可以接收第二个可选参数，一个表示通道的容量的整数。如果容量是0，make创建一个无缓冲通道。
```go
ch = make(chan int)         // 无缓冲通道
ch = make(chan int, 0)      // 无缓冲通道
ch = make(chan int, 3)      // 容量为3的缓冲通道
```

无缓冲通道的发送操作将会阻塞，直到另一个goroutine在对应的通道上执行接收操作，这时值传送完成，两个goroutine都可以继续执行。相反，如果先执行接收操作，接收方goroutine将阻塞，直到另一个goroutine在同一个通道上发送一个值。

使用无缓冲通道进行的通信会导致发送和接收goroutine同步化，因此，无缓冲通道也成为同步通道。

当一个值在无缓冲通道上传递时，接收值后发送方的goroutine才会再次被唤醒。

### 缓冲通道
创建一个可以容纳3个字符串的缓冲通道。
```go
ch = make(chan string, 3)
```

缓冲通道上的发送操作会在队列的尾部插入一个元素，接收操作从队列的头部移除一个元素。 如果通道满了，发送操作会阻塞所在的goroutine直到另一个goroutine对它进行操作来留可用的空间。如果通道是空的，执行接收操作的goroutine会阻塞，直到另一个goroutine在通道上发送数据。

可以使用内置函数cap获取缓冲通道的容量。
```go
fmt.Println(cap(ch))
```
可以使用内置函数len获取当前通道内的元素个数，不过由于在并发程序中这个信息会随着检索操作很快过时，不常用。
```go
fmt.Println(len(ch))
```

#### 示例

```go
func mirroredQuery() string {
responses := make(chan string, 3)
go func () { responses <- request("asia.gopl.io") }()
go func () { responses <- request("europe.gopl.io") }()
go func () { responses <- request("americas.gopl.io") }()
return <-responses // 返回最快的response。
}

func request(hostname string) (response string) {}
```

## 管道
通道可以用来连接goroutine，这样一个的输出是另一个的输入。这个叫管道（pipeline）。

没有一个直接的方式来判断是否通道已经关闭，但是有接收操作的一个变种，它产生两个结果：接收到的通道元素，以及一个布尔值（通常称为ok），它为true的时候代表接收成功，false表示当前的接收操作在一个关闭的并且读完的通道上。
```go
ch := make(chat int)
x, ok = <- ch 
```

程序结束时，关闭每一个通道不是必需的。只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。

通道可以通过垃圾回收器根据它是否可以访问来决定是否回收，而不是根据它是否关闭。

关闭一个已经关闭的通道会导致宕机。关闭通道还可以作为一个广播机制。

## 单向通道类型
单向通道，只允许发送或接收操作。

类型chan<- int是一个只能发送的通道，允许发送但不允许接收。

类型<-chan int是一个只能接收的int类型通道，允许接收但不能发送。

在任何赋值操作中将双向通道转换为单向通道都是允许的，但是反过来不行。 一旦有一个想chan<- int这样的单向通道，是没有办法通过它获取到引用同一个数据结构的chan int数据类型的。

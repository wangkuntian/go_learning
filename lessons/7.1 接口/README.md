# 接口

```go
type Phone interface {
call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
fmt.Println("I am iPhone, I can call you!")
}

var phone Phone

phone = new(NokiaPhone)
phone.call() // I am Nokia, I can call you!

phone = new(IPhone)
phone.call() // I am iPhone, I can call you!
```

# 接口类型

```go
var any interface{}
any = true
any = 12.13
any = "hello"
any = map[string]int{"one": 1}
any = new(bytes.Buffer)
```

空接口类型interface{}，空接口类型对其实现类型没有任何要求，所以可以把任意值赋给空接口类型。

# 接口值

一个接口类型的值（接口值）分为两个部分：一个具体类型和该类型的一个值。二者称为接口的动态类型和动态值。

下面四个语句中，变量w有三个不同的值（最初和最后是同一个值）。

```go
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
```

```go
var w io.Writer
```

接口的零值就是将它的动态类型和值都设置为nil。 一个接口值是否是nil取决于它的动态类型，所以现在这是一个nil接口值。可以用 w == nil或者w != nil来判断一个接口值是否是nil。 调用一个nil接口的任何方法都会导致崩溃。

```go
w.Write([]byte("hello")) //错误
```

```go
w = os.Stdout
```

将一个*os.File类型的值赋给w。这次赋值把一个具体类型隐式转换为一个接口类型，它与对应的显式转换io.Writer(os.Stdout)等价。

不管类型的转换是隐式的还是显式的，它都可以转换操作数的类型和值。

接口值的动态类型会设置为指针类型*os.File的类型描述符，它的动态值会设置为os.Stdout的副本，即一个指向代表程序的标准输出的os.File类型的指针。

调用该接口值的Write方法，会实际调用(*os.File).Write方法。

```go
w.Write([]byte("hello")) // hello
// 等价于
os.Stdout.Write([]byte("hello"))
```

```go
w = new(bytes.Buffer)
```

把一个*bytes.Buffer类型的值赋给了接口值。动态类型现在是\*bytes.Buffer，动态值现在则是一个指向新分配缓冲区的指针。

```go
w.Write([]byte("hello")) // 把hello写入bytes.Buffer
// 等价于
(*bytes.Buffer).Write([]byte("hello"))
```

```go
w = nil // 把nil赋值给了接口值
```


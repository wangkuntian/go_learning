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

# 实现接口
如果一个类型实现了一个接口要求的所有方法，那么这个类型实现了这个接口。

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

# 类型断言
类型断言是一个作用在接口值上的操作，写出来类似于x.(T)m，其中x是一个接口类型的表达式，而T是一个类型（称为断言类型）。类型断言会检查作为操作数的动态类型是否满足指定的断言类型。

如果断言类型T是一个具体类型，那么类型断言会检查x的动态类型是否就是T。如果检查成功，类型断言的结果就是x的动态值，类型当然就是T。

类型断言就是用来从它的操作数中把具体类型的值提取出来的操作。如果检查失败，那么操作崩溃。
```go
var w io.Writer
w = os.Stdout
f := w.(*os.File)       // 成功：f == os.Stdout
c := w.(*bytes.Buffer)  // 崩溃
```

如果断言类型T是一个接口类型，那么类型断言检查x的动态类型是否满足T。如果检查成功，动态值并没有提取出来，结果仍然是一个接口值，接口值的类型和值部分也没有变更，只是结果的类型为接口类型T。

类型断言是一个接口表达式，从一个接口类型变为拥有另外一套方法的接口类型（通常方法数量是增多），但保留了接口值中的动态类型和动态值部分。

```go

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p)) // convert int to ByteCounter
    return len(p), nil
}

var w io.Writer
w = os.Stdout
rw := w.(io.ReadWriter) // 成功： *os.File有Read和Write方法

w = new(ByteCounter)
rw = w.(io.ReadWriter) // 崩溃：*ByteCounter没Read方法
```

w和rw都有持有os.Stdout，于是所有对应的动态类型都是*os.File，当w作为io.Writer仅暴露了文件的Write方法，而rw还暴露了它的Read方法。

无论哪种类型作为断言类型，如果操作数是一个空接口值，类型断言都会失败。

如果类型断言出现在需要两个结果的赋值表达式中，那么断言不会在失败时崩溃，而是会多返回一个布尔类型的返回值来指示断言是否成功。
```go
var w io.Writer = os.Stdout
f, ok := w.(*os.File)       // 成功，ok, f == os.Stdout
b, ok := w.(*bytes.Buffer)  // 失败：!ok, b == nil
```
可以用if表达式的扩展形式。
```go
if f, ok :=w.(*os.File); ok {
	
}
```

当类型断言的操作数是一个变量时，返回值的名字可以和操作数变量名一致，原有的值就被新的返回值掩盖了。
```go
if w, ok := w.(*os.File); ok {
	
}
```


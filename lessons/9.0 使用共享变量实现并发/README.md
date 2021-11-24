# 并发安全
一个能在串行程序正确工作的函数，如果这个函数在并发调用时仍然能正确工作，那么这个函数就是并发安全（concurrency-safe）的。 如果一个类型的所有可访问方法和操作都是并发安全时，则它可以称为并发安全的类型。

函数并发调用时不工作的原因有很多，包括竞态、死锁、活锁（livelock）以及资源耗尽。

# 竞态
竞态是指在多个goroutine按某些交错顺序执行时程序无法给出正确的结果。

## 数据竞态
数据竞态发生于两个goroutine并发读写同一个变量并且至少其中一个是写入时。

## 避免竞态
1. 不修改变量，提前初始化数据。
2. 避免从多个goroutine访问同一个变量。不要通过共享内存来通信，而应该通过通信来共享内存。
3. 允许多个goroutine访问同一个变量，但同一时间只有一个goroutine可以访问。这种方法称为互斥机制。

## 竞态检测器
竞态检测器（race detector）-race命令加入到go build、go run、go test命令中即可使用该功能。

# 互斥锁：sync.Mutex
```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
```

# 读写互斥锁：sync.RWMutex
```go
var mu sync.RWMutex
mu.RLock()              // 读锁
defer mu.RUnlock()
mu.WLock()              // 写锁
defer mu.WUnlock()
```

RLock仅可用于在临近区域内对共享变量无写操作的情形。
仅在绝大部分goroutine都在获取读锁且锁竞争比较激烈时（goroutine一般都需要等待后才能获取到锁），RWMutex才有优势。因为RWMutex需要更复杂的内部簿记工作，所以在竞争不激烈时它比普通的互斥锁慢。

# 内存同步
同步不仅仅涉及多个goroutine的执行顺序问题，同步还会影响到内存。
现代的计算机一般都会有多个处理器，每个处理器都有内存的本地缓存。为了提高效率，对内存的写入是缓存在每个处理器中的，只在必要时才刷回内存。甚至刷回内存的顺序都有可能与goroutine的写入顺序不一致。像通道通信或者互斥锁操作这样的同步原语都会导致处理器把累计的写操作刷回内存并提交，所以这个时刻之前goroutine的执行结构就保证了对运行在其他处理器的goroutine可见。

考虑如下代码片段可能的输出：
```go
var x, y int
go func() {
	x = 1                       // A1
	fmt.Print("y:", y, " ")     // A2
}()
go func() {
    y = 1                       // B1
    fmt.Print("x:", x, " ")     // B2
}()
```
由于这个两个goroutine并发且在没有使用互斥锁的情况下共享变量，所以这里会有数据竞态。可能会得到如下的结果之一。
```go
y:0 x:1
x:0 y:1
x:1 y:1
y:1 x:1
```
第四种结果可以由A1，B2，A2，B2或者B1，A1，A2，B2这样的顺序来产生。但是，程序如果出现下面两种输出就出人意料了。
```go
x:0 y:0
y:0 x:0
```
但在某种特定的编译器、CPU或者其他情况下，这些确实可能发生。

在单个goroutine内，每个语句的效果保证按照执行的顺序发生。也就是说，goroutine是串行一致（顺序一致）的。但在缺乏使用通道或者互斥量来显式同步的情况下，并不能保证所有的goroutine看到的事情顺序都是一致的。
尽管goroutine A肯定能在读取y之前观察到x=1的效果，但是它不一定能观察到goroutine B对y的写入效果，所有A可能会输出y的一个过期值。

尽管很容易把并发简单理解为多个goroutine中语句的某种交错执行方式，但正如上面的例子所显示的，这并不是一个现代编译器和CPU的工作方式。因为赋值和Print对应不同的变量，所以编译器就可能会认为两个语句的执行顺序不会影响结果，然后就交换了这两个语句的执行顺序。CPU也有类似的问题，如果两个goroutine在不同的CPU上执行，每个CPU都有自己的缓存，那么一个goroutine的写入操作在同步到内存之前对另外一个goroutine的Print语句是不可见的。

这些并发问题都可以通过采用简单、成熟的模式来避免，即在可能的情况下，把变量限制到单个goroutine中，对于其他变量，使用互斥锁。

# 延迟初始化：sync.Once
示例
```go
var icons map[string]image.Image

func loadIcon(name string) image.Image {
    return icons[name]
}

// 并发不安全
func Icon(name string) image.Image {
    if icons == nil {
        loadIcons()
    }
    return icons[name]
}

func loadIcons() {
    icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
        "hearts.png":   loadIcon("hearts.png"),
        "diamonds.png": loadIcon("diamonds.png"),
        "clubs.png":    loadIcon("clubs.png"),
    }
}
```
上面的这例子，在并发调用Icon时这个模式是不安全的。

在缺乏显式同步的情况下，编译器和CPU在能保证每个goroutine都满足串行一致性的基础上可以自由地重拍访问内存的顺序。loadIcons一个可能的语句重排结果如下所示。它在填充数据之前吧一个空map赋给icons。
```go
func loadIcons(){
	icons = make(map[string]image.Image)
	icons["spades.png"] = loadIcon("spades.png")
	icons["hearts.png"] = loadIcon("hearts.png")
	icons["diamonds.png"] = loadIcon("diamonds.png")
	icons["clubs.png"] = loadIcon("clubs.png")
}
```
因此，一个goroutine发现icons不是nil并不意味着变量初始化肯定已经完成。
保证所有goroutine都能观察到loadIcons效果最简单的正确方法就是用一个互斥锁来做同步。
```go
var mu sync.Mutex   // 保护icons
var icons map[string]image.Image

// 并发安全
func Icon(name string) image.Image {
	mu.Lock()
	defer mu.Unlock()
    if icons == nil {
        loadIcons()
    }
    return icons[name]
}
```
采用互斥锁访问icons的额外代价是两个goroutine不能并发访问这个变量，即使变量在已经完全完成初始化且不再更改的情况下，也会造成这个后果。使用一个可以并发读的锁就可以改善这个问题。
```go
var mu sync.RWMutex   // 保护icons
var icons map[string]image.Image

// 并发安全
func Icon(name string) image.Image {
    mu.RLock()
    if icons == nil {
		icon := icons[name]
        mu.RUnlock()
		return icon
    }
	mu.RUlock()
	mu.Lock()
	if icons == nil {
		loadIcons()
    }
	icon := icons[name]
	mu.Unlock()
    return icon
}
```
这里有两个临界区域。goroutine首先获得一个读锁，查阅map，然后释放这个读锁。如果条目能够找到，就返回它。如果条目没有找到，goroutine再获取一个写锁。由于不先释放一个共享锁就无法直接把它升级到互斥锁，为了避免在过渡期其他goroutine已经初始化了icons，所以必须重新检查nil值。

sync包提供了针对一次性初始化问题的特化解决方案：sync.Once。从概念上讲，Once包含一个布尔变量和一个互斥量，布尔变量记录初始化是否已经完成，互斥量则负责保护这个布尔变量和客户端的数据结构。
Once的唯一方法Do以初始化函数作为它的参数。
```go
var loadIconsOnce sync.Once
// Icon 并发安全
func Icon(name string) image.Image {
    loadIconsOnce.Do(loadIcons)
    return icons[name]
}
```
每次调用Do(loadIcons)时，会先锁定互斥量并检查里边的布尔变量。在第一次调用时，这个布尔变量为假，Do会调用loadIcons然后把变量设置为真。后续的调用相当于空操作，只是通过互斥量的同步来保证loadIcons对内存产生的效果对所有的goroutine可见。以这种方式来使用sync.Once，可以避免变量在正确构造之前就被其他goroutine分享。

# goroutine和线程
goroutine和线程的区别

## 栈内存
每个OS线程都有一个固定大小的栈内存（通常为2MB）。栈内存区域用于保存其他函数调用期间那些正在执行或者临时暂停的函数中的局部变量。

对于一个小的goroutine，2MB的栈无疑是个巨大的浪费。同时在Go程序中一次性创建十万左右的goroutine并不罕见，这时栈内存的开销是相当巨大的。另外，对于最复杂和深度递归的函数，固定大小的栈始终不够大。

一个goroutine在生命周期开始只有一个很小的栈内存，典型情况下为2KB。 与OS线程类似，goroutine的栈也用于存放那些正在执行或临时暂停的函数中的局部变量。与OS线程不同的是，goroutine的栈不是固定大小的，它可以按需增大或缩小。

## 调度
OS线程由OS内核调度。每隔几毫秒，一个硬件时钟中断发送到CPU，CPU调用一个叫调度器的内核函数。这个函数暂停当前正在运行的线程，
把它的寄存器信息保存到内存，查看线程列表并决定接下来运行哪一个线程，再从内存恢复线程的注册表信息，最后继续执行选中的线程。

因为OS线程由内核来调度，所以控制权从一个线程到另一个线程需要一个完整的上下文切换（context switch）：即保存一个线程的状态到内存，
再恢复另一个线程的状态，最后更新调度器的数据结构。考虑这个操作涉及的内存局域性以及涉及的内存访问数量，还有访问内存所需的CPU周期数量的增加，
这个操作会很慢。

Go运行时包含一个自己的调度器，这个调度器使用一个称为m:n调度技术（它可以复用/调度m个goroutine到n个OS线程）。Go调度器与内核调度器的工作类似，
但是Go调度器只需关心单个Go程序的goroutine调度问题。

与操作系统的线程调度器不同的是，Go调度器不是由硬件时钟来定期触发的，而是由特定的Go语言结构来触发的。

比如当一个goroutine调用time.Sleep或被通道阻塞或对互斥量操作时，调度器就会将这个goroutine设为休眠模式，
并运行其他goroutine直到前一个可重新唤醒为止。因为它不需要切换到内核语境，所以调度一个goroutine比调度一个线程成本低很多。

## GOMAXPROCS
Go调度器使用GOMAXPROCS参数来确定需要使用多个OS线程来同时执行Go代码。默认值是机器上的逻辑CPU数量，所以在一个有着8核CPU的机器上，
调度器会把Go代码同时调度到8个OS线程上。（GOMAXPROCS是m:n调度中的n）

正在休眠或者正在被通道通信阻塞的goroutine不需要占用线程。

可以使用GOMAXPROCS环境变量或者runtime.GOMAXPROCS函数来显式控制这个参数。

## 标识
在大部分支持多线程的操作系统和编程语言里，当前线程都有一个独特的标识，它通常可以取一个整数或者指针。这个特性可以让我们轻松构建一个线程的局部存储，
它本质上就是一个全局的map，以线程的标识作为键，这样每个线程都可以独立地用这个map存储和获取值，而不受其他线程干扰。

goroutine没有可供程序员访问的标识，因为线程局部存储有一种被滥用的倾向。





# go test工具
go test命令是Go语言包的测试驱动程序。在一个包目录中，以_test.go结尾的文件不是go build命令编译的目标，而是go test编译的目标。

在*_test.go文件中，三种函数需要特殊对待，功能测试函数、基准测试函数和实例函数。

功能测试函数时以Test前缀命名的函数，用来检测一些程序逻辑的正确性，go test运行岑石函数，并且报告是PASS还是FAIL。
基准测试函数以Benchmark开头，用来测试某些操作的性能，go test汇报操作的平均执行时间。
实例函数以Example开头，用来提供机器检查过的文档。

go test工具扫描*_test.go来寻找特殊函数，并生成一个临时的main包来调用它们，然后编译运行，并汇报结果，最后清空临时文件。

# 功能测试函数
每个测试文件必须导入testing包。
```go
func TestName(t *testing.T){
	
}
```
功能测试函数必须以Test开头，可选的后缀名称必须以大写字母开头。
```go
func TestSin(t *testing.T) { }
func TestCos(t *testing.T) { }
func TestLog(t *testing.T) { }
```
参数t提供了汇报测试失败和日志记录的功能。

go test -v可以输出包中每个测试用例的名称和执行的时间。
go test -run的参数是一个正则表达式，它可以使得go test只运行那些测试函数名称匹配给定模式的函数。

# 基准测试函数
基准测试就是在一定工作负载之下检测程序西能的一种方法。

在Go里面，基准测试函数看上去像一个测试函数，但是前缀是Benchmark并且拥有一个\*testing.B参数用来提供大多数和\*testing.T相同的方法，额外增加了一些与性能检测相关方法。它还提供了一个整型成员N，用来指定被检测操作的执行次数。
```go
import "testing"

func BenchmarkIsPalindrome(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPalindrome("A man, a plan, a canal: Panama")
    }
}
```
go test -bench执行基准测试。
```shell
go test -bench=.
goos: darwin
goarch: amd64
pkg: world
cpu: Intel(R) Core(TM) i5-6267U CPU @ 2.90GHz
BenchmarkIsPalindrome-4          6974612               167.4 ns/op
PASS
ok      world   1.350s
```

-bench的参数指定了要运行的基准测试，它是一个匹配Benchmark函数名称的正则表达式，它的默认值不匹配任何函数。模式"."匹配word包中所有的基准测试函数，因为这里只有一个基准测试函数，所以这个和指定-bench=IsPalindrome效果是一样的。

基准测试名称的数字后缀4表示GOMAXPROCS的值，这个对于并发基准测试很重要。

基准测试报告显示每次IsPalindrome调用0.1674ms，这个是6974612次调用的平均值。因为基准测试运行器开始时并不清楚这个操作的耗时长短，所以开始的时候它使用了比较小的N值来做检测，然后为了检测稳定的运行时间，推断出足够大的N值。

使用基准测试函数来实现循环而不是在测试驱动程序中调用代码的原因是，在基准测试函数中在循环外面可以执行一些必要的初始化代码并且这段时间不会被添加到每次迭代的时间中。

对两种不同的算法使用相同的输入，在重要的或具有代表性的工作负载下，进行基准测试通常可以显示出每个算法的优点和缺点。

# 示例函数
示例函数的前缀是Example。它既没有参数也不会返回结果。
```go
func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
}
```

示例函数有三个目的。
1. 首要目的是作为文档，比起乏味的描述，举一个好的例子是描述库函数功能最简介直观的方式。
2. 示例函数是可以通过go test运行的可执行测试。
3. 示例函数提供了手动实验的代码。
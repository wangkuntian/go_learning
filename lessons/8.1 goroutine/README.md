
# goroutine
在Go里，每一个并发执行的活动称为goroutine。
```go
f()         // 调用f()；等待它返回
go f()      // 新建一个调用f()的goroutine，不用等待。
```

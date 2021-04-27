
# 互斥锁
```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
```

在单个goroutine内，每个语句的效果保证按照执行的顺序发生。
goroutine是串行一致（顺序一致）的。但在缺乏使用通道或者互斥量来显式同步的情况下，并不能保证所有的goroutine看到的事情顺序都是一致的。
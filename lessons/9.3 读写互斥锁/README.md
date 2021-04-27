
# 读写互斥锁
```go
var mu sync.RWMutex
mu.RLock()              // 读锁
defer mu.RUnlock()
mu.WLock()              // 写锁
defer mu.WUnlock()
```

RLock仅可用于在临近区域内对共享变量无写操作的情形。
仅在绝大部分goroutine都在获取读锁且锁竞争比较激烈时（goroutine一般都需要等待后才能获取到锁），RWMutex才有优势。因为RWMutex需要更复杂的内部簿记工作，所以在竞争不激烈时它比普通的互斥锁慢。


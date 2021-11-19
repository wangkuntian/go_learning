
# Map
使用内置函数make声明map。
```go
ages := make(map[string]int) // 创建一个从string到int的map
ages["alice"] = 20
ages["charlie"] = 24
```
相当于。
```go
ages := map[string]int {
	"alice": 20,
	"charlie": 24,
}
```

使用delete函数从map中移除元素。
```go
delete(ages, "alice")
```

即使key不在map中，上述操作也是安全的，当key不存在时，就返回值类型的零值。
```go
ages["bob"] = ages["bob"] + 1 // 0 + 1
```
map元素不是一个变量，无法获取它的地址。

# 遍历
```go
for name, age := range ages {
	fmt.Printf("%s \t %d \n", name, age)
}
```

map中元素的迭代顺序是不固定的。

map的零值是nil。

判断两个map是否拥有相同的键和值。
```go
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
```


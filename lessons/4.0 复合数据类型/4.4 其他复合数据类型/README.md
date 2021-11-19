
# JSON
将Go的数据结构转换为JSON称为marshal。
json.Marshal来可以将Go的数据结构转换为JSON。
json.MarshalIndent可以输出整齐格式化过的结果。
json.Unmarshal可以将JSON字符串解码为Go数据结构。
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"},
	},
	{
		Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
	},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
```

marshal使用Go结构体成员的名称作为JSON对象里字段的名称（反射）。
只有可导出的成员可以装换为JSON字段。所以将GO结构体中所有成员都定义为首字母大写。

成员标签定义是结构体成员在编译期间关联的一些元信息。
```go
Year   int  `json:"released"`
Color  bool `json:"color,omitempty"`
```

omitempty表示如果这个成员值是零值或者为空，则不输出这个成员到JSON中。

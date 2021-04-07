
# 整型

## 有符号
int int8 int16 int32 int64

int和uint大小相同，都是32或者64位。
rune和int可以互换使用
byte可以和int8互换使用，byte强调的是原始数据。

有符号数以补码表示，保留最高位作为符号位，n位数的取值范围是-2^(n - 1) ~ 2^(n - 1) - 1 

## 无符号
uint uint8 uint16 uint32 uint64
无符号整数由全部位构成其非负值，范围是0 ~ 2^n - 1

## 二目运算符
\*   /   %   <<（左移）  >>（右移）  &（与）   &^（位清空）
z=x&^y中，若y的某位是1，则z的对应位等于0，否则，它就等于x的对应位。
\+   -   |（或）   ^（异或）
\==  !=  <   <=  >   >=
\&&
\||

# 浮点型
## float32
最大值math.MaxFloat32，大约为3.4e38，最小值正浮点数约为1.4e-45。

## float64
最大值math.MaxFloat64，大约为1.8e308，最小值正浮点数约为4.9e-324。

绝大多数情况下，优先选用float64。

# 复数
complex64 complex128，分别由float32和float64构成。
```go
var x complex128 = complex(1, 2)    // 1 + 2i
real(x)                             // 1
imag(y)                             // 2
```

# 布尔型
true false

# 字符串
utf-8编码

```go
import "unicode/utf8"

s := "Hello, 世界"
fmt.Println(len(s))                 // "13"
fmt.Println(utf8.RuneInString(s))   // "9"

```

字符串转换
```go
strconv.Itoa(123)                   // "123"
strconv.Atoi("123")                 // 123 整型
strconv.ParseInt("123", 10, 64)     // 十进制，最长为64位
```


# 常量
布尔型、字符串或数字（整型和浮点型）

## 常量生成器iota
```go
const(
    _ = 1 << ( 10 * iota)
    KiB                     // 1024
    MiB                     // 1048576
    GiB
    TiB
    PiB
    EiB
    ZiB
    YiB
)
```

## 无类型常量
无类型布尔
无类型整数
无类型浮点数
无类型复数
无类型字符串
无类型文字符号

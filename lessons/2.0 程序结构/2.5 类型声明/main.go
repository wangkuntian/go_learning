package main

// Celsius 摄氏度
type Celsius float64

// Fahrenheit 华氏度
type Fahrenheit float64

const (
	// AbsouluteZeroC 绝对0度
	AbsouluteZeroC Celsius = -273.15
	// FreezingC 水的冰点
	FreezingC Celsius = 0
	// BoilingC 水的沸点
	BoilingC Celsius = 100
)

// CToF 摄氏度转华氏度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC 华氏度转摄氏度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	var c Celsius = 100
	f := CToF(c)
	FToC(f)
}

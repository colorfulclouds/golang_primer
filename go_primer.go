package main

import (
	"fmt"
)

func main() {
	// var aa int
	// aa = 10
	// // 与下面的写法是一样的
	// v := 10

	// // 多重赋值
	// ab, ba := 10, 20
	// const az = 10
	// const a, b = 10, 3.14

	// 两种写法是一样的
	// 常量不需要使用冒号等号运算符
	const c = 45
	// const c int = 45
	fmt.Println(c)

	// _ 下划线 是匿名变量

	var a int = 1
	var b float64 = 2.0
	// 与下面的写法是一样的
	var (
		a int     = 1
		b float64 = 2.0
	)

	var (
		a = 1 // 类型自动推导
		b = 2.0
	)

	const i int = 10
	const j float64 = 3.14
	//这些写法都是一样的
	const i = 10
	const j = 3.14
	const (
		i int     = 10
		j float64 = 3.14
	)

	const (
		i = 10
		j = 3.14
	)

	// iota的学习
	// iota 常量自动生成器 每个一行 自动累加1
	// 只给常量赋值使用
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	//iota遇到const 重置为0
	const d = iota // d=0
	// 可以只写一个iota
	const (
		a = iota // 0
		b        // 1
		c        // 2
	)
	// 如果在一行 所有的值一样
	const (
		i          = iota             // 0
		j1, j2, j3 = iota, iota, iota // 1
		k          = iota             // 2
	)
	byte    // 1字节也是uint8的别名 存ascii码 就是C语言中的char类型
	rune    // 4字节也是uint32的别名 存unicode码
	uintptr // 4 or 8 存储指针的uint32 uint64整数

	pi := 3.14 // 类型推导时 推导为float64

	fushu := 1 + 3.14i //自动类型推导为 complex128
	// 使用内建函数 获得实部与虚部
	fmt.Printf("real:", real(fushu), "imag:", imag(fushu))
}

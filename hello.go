package hello

import "fmt"

// 一个用于演示如何发布自定义包的示例项目。

// SayHi 一个打招呼的函数
func SayHi(name string) {
	fmt.Printf("你好%s, 我是jokereven. 很高兴认识你~\n", name)
}

// SayBye 一个打招呼的函数
func SayBye(name string) {
	fmt.Printf("你好%s, 我是jokereven. 期待下次相见~\n", name)
}

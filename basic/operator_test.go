package main 

import (
	"fmt"
	"testing"
)

var (
	a = 1
	b = 2
	c = 3
)

func TestOperator(t *testing.T) {
	var d int // 默认初始化为0
	fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)
	// 算数运算
	// + - * / % ++ --
	d = b + a 
	fmt.Println("+ d = ", d)
	d = b - a 
	fmt.Println("- d = ", d)
	d = b * c
	fmt.Println("* d = ", d)
	d = c % b
	fmt.Println("% d = ", d)
	d-- 
	fmt.Println("-- d = ", d)
	d++ 
	fmt.Println("++ d = ", d)

	// 关系运算
	// == != >= <= > <
	
	// 赋值运算
	// = += -= *= /= %=

	// 逻辑运算
	// && || !

	// 位运算符
	// >> << ^ & |
}
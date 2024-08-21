package main

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	const PI = 3.14
	fmt.Println("PI_first = ", PI)
	// PI = 14
	// fmt.Println("PI_after = ", PI)
	// fmt.Println(*PI)

	var a int = 1
	fmt.Println("a = ", a)
	fmt.Println(&a)
	b := &a
	fmt.Println(b)
	a = 2
	fmt.Printf("b = %v\n", *b)
	fmt.Printf("b = %T", b)
	// fmt.Println(a)
}
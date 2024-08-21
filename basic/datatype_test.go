package main

import (
	"fmt"
	"testing"
	"sync"
)
var global int = 1
// func init() {
// 	fmt.Println("init function1")
// }

// func init() {
// 	fmt.Println("init function2")
// }

// func init() {
// 	global = 2
// 	fmt.Println(global)
// }

// func init() {
// 	global = 3
// 	fmt.Println(global)
// }


func TestBasicDateType(t *testing.T) {
	fmt.Println("******** global = ", global)
	// 整型（包括字符型）int8 int32 int64
	// byte=uint8 rune=int32
	var str string = "中国people" 
	for _, value :=range str {
		fmt.Println(value)  // 容纳更多字符集
		fmt.Printf("value = %T\n", value)
	}
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[i] = %v\n", string(str[i])) // 按照ASCII码来识别，对于中文等特殊字符无法识别
		fmt.Printf("str[i] = %T\n", str[i]) 
	}

	// 浮点型
	// float32 float64

	// 布尔类型
	var choice1 bool = false
	// var choice2 bool = 0  // golang中不可使用0，1表示布尔值
	if !choice1 {
		fmt.Println("choice1")
	}
	// if choice2 {
	// 	fmt.Println("choice2")
	// }

	// 字符串
	// 复数
}

// 类型转换 无法隐式转换
func TestChangeType(t *testing.T) {
	fmt.Println("============")
	var a int = 1
	b := float32(a)
	fmt.Println(b)
	fmt.Printf("str = %T\n", b)
}
// 类型断言
func TestTypeAssertion(t *testing.T) {
	var a interface{} = 10
	_, v := a.(int)
	if v {
		fmt.Println("a is type int")
	} else {
		fmt.Println("a is not type int")
	}
}
// 类型别名
// 类型定义
func TestDateType(t *testing.T) {
	// 数组、切片、map、channel、函数、接口、结构体、指针
	// Array
	var arr [8]int = [8]int{1, 2, 3, 4}
	fmt.Println(arr)

	var arr2 = new([8]int)
	fmt.Printf("arr2 = %T\n", arr2)
	fmt.Println("arr2 = ", arr2)
	// golang中数组长度不可变
	// 改
	arr[3] = 100
	fmt.Println(arr)
	// 查 能查找指定下标元素或者指定元素的下标

	// Slice
	// var arr3 = make([8]int)
	// fmt.Printf("arr3 = %T\n", arr3)
	// fmt.Println("arr3 = ", arr3)
	var slice []int = []int{1, 2, 3, 4}
	fmt.Println("slice = ", slice)
	fmt.Printf("slice = %T\n", slice)
	var slice1 = make([]int, 8)
	fmt.Println(slice1)
	fmt.Printf("slice1 = %T\n", slice1)

	var slice2 = arr[:1]
	fmt.Println("before = ", slice2)
	slice2[0] = 100
	fmt.Println("after = ", slice2)
	fmt.Println("arr = ", arr) // slice2是arr的引用，所以可以改变arr

	// 增
	slice4 := []int{101, 102, 103}
	slice3 := append(slice1, 10, 2, 3)
	fmt.Println(slice3)
	slice3 = append(slice3, slice4...)
	fmt.Println(slice3)
	// 删
	slice3 = slice3[:3]
	fmt.Println(slice3)
	// 改或者查，与array操作一致， 注意如果切片是数组分割生成的，那更改切片会影响原数组的元素，因为切片是数组的引用
	// nil或者空slice

	// Map
	var map1 = map[string]int{"a":1, "b":2}
	fmt.Println(map1["a"])
	fmt.Println(map1["c"])
	map2 := make(map[string]int)
	fmt.Println(map2)
	// 增或者改，操作差不多
	map1["c"] = 3
	fmt.Println(map1)
	// 删
	delete(map1, "c")
	delete(map1, "d") // 即使元素不存在，执行delete函数也不会报错
	// delete(map1, 3) 删除不可使用value值
	fmt.Println(map1)
	// 查
	// 根据key来查找value
	// v := map1["a"] 一个返回值时，返回查找key对应的value
	k, v := map1["a"] // 两个返回值时候返回查找key对应的value 和 是否存在该值布尔值； 如果不存在则返回零值和false
	fmt.Println(k, v)

	// map 是无序（底层实现一个函数打乱key），且线程不安全的
	var map3 = make(map[int]int)
	for i := 1; i < 100; i++ {
		map3[i] = i
	}
	for k, value := range map3 {
		fmt.Println(k , value)
	}
	
	// 指针 

	// channel 

	// interface

	// // 针对函数和结构体，了解方法定义和使用
}

var wg1 sync.WaitGroup
//函数: 闭包函数，回调函数callback 匿名函数 内置函数 defer init函数
func TestFunction(t *testing.T) {
	j := 10
	for i := 0 ; i < 10; i++ {
		j--
		wg1.Add(1)
		go func(i int) {
			fmt.Printf("i = %v, j = %v\n", i, j)
			wg1.Done()
		}(i)
		go printNum(j, func() {
			fmt.Println("this is a test")
		})
	}

	// 匿名函数 ,可以直接执行
	func() {
		fmt.Println("Anonymous Function")
	}()
	wg1.Wait()
}

func printNum(i int, callback func()) {
	fmt.Println("num = ", i)	
	callback()
}

// defer 
func TestDefer(t *testing.T) {
	var a = 100
	defer func() {
		fmt.Println("the end of code")
	}()
	defer func() {
		return
		fmt.Println("defer a = ", a)
	}()
	a = 2
	fmt.Println("a = ", a)
}

// interface
type Skill interface {
	Fly() bool
}

// 结构体 嵌套结构体
type People struct{
	age int 
	name string
}

func (p *People) Fly() bool {
	fmt.Println("People can fly!")
	return true
}

func TestStruct(t *testing.T) {
	var p1 = People{
		age: 20,
		name: "mike",
	}
	fmt.Println(p1.age)
	res := p1.Fly()
	fmt.Println("res = ", res)
}


// channel
func TestChannel(t *testing.T) {
	var channel chan int
	channel = make(chan int)
	fmt.Printf("channel = %T\n", channel)

	go func(channel chan int) {
		channel <- 1
	}(channel)
	res := <-channel
	fmt.Println("res = ", res) // 将channel接受的值赋值给res
	// fmt.Println(<-channel) 将channel接受的值直接忽略，无打印

	channel2 := make(chan int, 8)
	channel2 <- 1
	channel2 <- 2
	channel2 <- 3
	channel2 <- 4
	fmt.Println("len(channel2) = ", len(channel2))
	fmt.Println("cap(channel2) = ", cap(channel2))
}

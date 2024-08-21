package main 

import (
	"fmt"
	_ "time"
	"sync"
)

func main() {
	// i := 0
	// for ; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 	}(i) // 闭包
	// }
	// fmt.Println("i = ", i)
	sync_exec()
	// time.Sleep(time.Duration(2)*time.Second) // 使用time.Sleep 暂停程度5秒
}


func sync_exec() {
	i := 0
	var wg sync.WaitGroup
	wg.Add(10)
	for ; i < 10; i++ {
		// wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i) // 闭包
	}
	wg.Wait()  // 使用sync.WaitGroup 
}
package main

import (
	"testing"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex
var once sync.Once

func TestMutex(t *testing.T) {
	var num int = 1
	fmt.Println("num begin =", num)
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				fmt.Println("Only do once")
			})
			mu.Lock()
			num = num + 1
			time.Sleep(time.Duration(1)*time.Second)
			fmt.Printf("num = %v\n", num)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
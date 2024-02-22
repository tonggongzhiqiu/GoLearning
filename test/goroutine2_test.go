package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func myGoroutine2(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Printf("myGoroutine %s\n", name)
		time.Sleep(2)
	}
}

func TestGoroutine2(t *testing.T) {
	//var wg *sync.WaitGroup // 这么写报错因为没有实例化
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go myGoroutine2("goroutine1", wg)
	go myGoroutine2("goroutine2", wg)

	wg.Wait()
}

func TestGorutine22(t *testing.T) {
	var wg sync.WaitGroup
	fmt.Printf("wg %v\n", wg)
	wg.Add(2)

	go myGoroutine2("goroutine11", &wg)
	go myGoroutine2("goroutine22", &wg)

	wg.Wait()
}

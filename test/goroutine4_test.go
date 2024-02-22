package test

import (
	"fmt"
	"testing"
)

func TestGoroutine4(t *testing.T) {
	case1 := make(chan int)
	case2 := make(chan int)
	close(case1)
	close(case2)

	select {
	// 关闭 channel 时，接收方会收到 false value，所以这里可以接收到内容
	case <-case1:
		fmt.Printf("case1")
	case <-case2:
		fmt.Println("case2")
	default:
		fmt.Println("default")
	}
}

package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func func1(ctx context.Context) {
	hctx, hcancel := context.WithTimeout(ctx, time.Second*4)
	defer hcancel()

	resp := make(chan struct{}, 1)
	// 处理逻辑
	go func() {
		// 处理耗时
		time.Sleep(time.Second * 10)
		resp <- struct{}{}
	}()

	// 超时机制
	select {
	//	case <-ctx.Done():
	//		fmt.Println("ctx timeout")
	//		fmt.Println(ctx.Err())
	case <-hctx.Done():
		fmt.Println("hctx timeout")
		fmt.Println(hctx.Err())
	case v := <-resp:
		fmt.Println("test2 function handle done")
		fmt.Printf("result: %v\n", v)
	}
	fmt.Println("test2 finish")
	return

}

// TestContext1 通过 context 的超时机制，定时取消协程的进行
func TestContext1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	func1(ctx)
}

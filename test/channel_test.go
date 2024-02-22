package test

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	var chanX chan int
	fmt.Printf("chanX %v\n", chanX)
	// 对于包含引用类型（如指针、切片、映射、通道和接口）的结构体，使用var声明时会被初始化为nil。这是因为引用类型的零值就是nil
	var ptr *int // 指针变量初始化为nil
	fmt.Printf("ptr %v\n", ptr)
	var s []int // 切片变量初始化为nil
	fmt.Printf("s %v\n", s)
	var m map[string]int // 映射变量初始化为nil
	fmt.Printf("m %v\n", m)
	var i interface{} // 接口变量初始化为nil
	fmt.Printf("i %v\n", i)
	//var f func() // 函数变量初始化为nil (暂时不知道怎么打印)
	//fmt.Printf("%+v\n", f)
	//fmt.Printf("fff %+v\n", f)
}

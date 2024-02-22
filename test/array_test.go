package test

import (
	"fmt"
	"testing"
)

func TestArray1(t *testing.T) {
	array := [3]int{1, 3, 4}

	test(array)
	fmt.Printf("array out %p \n", &array)
}

func TestArray2(t *testing.T) {
	array1 := [3]int{1, 3, 4}
	var array2 [3]int

	array2 = array1
	array1[0] = 10

	fmt.Println(array1)
	fmt.Println(array2)
}

func test(array [3]int) {
	fmt.Printf("array in %p \n", &array)
}

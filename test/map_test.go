package test

import (
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m := make(map[int]int)

	go func() {
		ReadMap(m)
	}()

	go func() {
		WriteMap(m)
	}()

	go func() {
		DeleteMap(m)
	}()

	time.Sleep(10 * time.Second)
}

func ReadMap(m map[int]int) {
	for i := 0; i < nums; i++ {
		go func(i int) {
			_, _ = m[i]
		}(i)
	}
}

func WriteMap(m map[int]int) {
	for i := 0; i < nums; i++ {
		go func(i int) {
			m[i] = i
		}(i)
	}
}

func DeleteMap(m map[int]int) {
	for i := 0; i < nums; i++ {
		go func(i int) {
			delete(m, i)
		}(i)
	}
}

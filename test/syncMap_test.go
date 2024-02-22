package test

import (
	"sync"
	"testing"
	"time"
)

var nums = 1000

func TestSyncMap(t *testing.T) {
	m := &sync.Map{}

	go func() {
		Read(m)
	}()

	go func() {
		Write(m)
	}()
	go func() {
		Write(m)
	}()
	go func() {
		Write(m)
	}()

	go func() {
		Delete(m)
	}()

	time.Sleep(10 * time.Second)
}

func Read(m *sync.Map) {
	for i := 0; i < nums; i++ {
		go func(i int) {
			m.Load(i)
		}(i)
	}
}

func Write(m *sync.Map) {
	for i := 0; i < nums; i++ {
		go func(i int) {
			m.Store(i, i)
		}(i)
	}
}

func Delete(m *sync.Map) {
	for i := 0; i < nums; i++ {
		go func(i int) {
			m.Delete(i)
		}(i)
	}
}

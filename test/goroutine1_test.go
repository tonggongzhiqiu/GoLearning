package test

import "testing"

func myGoroutine(t *testing.T) {
	//time.Sleep(2)
	t.Logf("my Groutine...")
}

func TestGoroutine(t *testing.T) {
	go myGoroutine(t)
	t.Logf("test end!!")
}

package test

import "sync"

var mu sync.Mutex
var count = 0

func add() {
	mu.Lock()
	count++
	mu.Unlock()
}

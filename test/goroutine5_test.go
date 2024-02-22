package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine5(t *testing.T) {
	catCh := make(chan struct{})
	dogCh := make(chan struct{})
	fishCh := make(chan struct{})

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			fmt.Println("cat")
			close(catCh)
			break
		}
	}()

	go func() {
		<-catCh
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			fmt.Println("dog")
			close(dogCh)
			break
		}
	}()

	go func() {
		<-dogCh
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			fmt.Println("fish")
			close(fishCh)
			break
		}
	}()

	<-fishCh
}

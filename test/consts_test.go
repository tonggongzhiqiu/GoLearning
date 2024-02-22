package test

import (
	"fmt"
	"testing"
)

type TypeA string

const (
	Car   TypeA = "car"
	Horae TypeA = "horae"
)

func TestConst(t *testing.T) {
	a := TypeA("car")
	fmt.Printf(string(a))
}

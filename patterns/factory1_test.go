package patterns

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi, my name is %s", p.Name)
}

// 简单工厂，返回 struct
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

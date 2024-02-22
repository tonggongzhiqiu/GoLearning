package patterns

import (
	"fmt"
)

// IPerson 抽象工厂，返回接口
// person 结构体是不可以在外部使用的
type IPerson interface {
	GreetName()
}

type person struct {
	name string
	age  int
}

func NewPerson2(name string, age int) IPerson {
	return &person{
		name: name,
		age:  age,
	}
}

func (p *person) GreetName() {
	fmt.Printf("Hi! My name is %s", p.name)
}

//
//func TestGreetName(t *testing.T) {
//	person := NewPerson2("zhangsan", 10)
//	person.GreetName()
//}

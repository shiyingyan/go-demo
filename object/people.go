package object

import (
	"fmt"
	"reflect"
)

type People struct {
	name   string
	age    uint8
	height float32
	weight float32
}

func (p *People) Newyear() *People {
	x := p
	fmt.Println(x, reflect.TypeOf(x), p, reflect.TypeOf(p))
	p.age += 1
	return p
}

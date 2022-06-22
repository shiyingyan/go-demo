package greetings

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello,%s welcome!", name)
}

type Gender uint

const (
	Male Gender = iota + 1
	Female
	//Male Gender = "male"
	//Female = "female"
)

type People struct {
	name   string
	age    uint8
	gender Gender
}

type Option func(people *People)

func WithGender(gender Gender) Option {
	return func(people *People) {
		people.gender = gender
	}
}

func WithAge(age uint8) Option {
	return func(people *People) {
		people.age = age
	}
}

func New(name string, opts ...Option) *People {
	p := People{name: name, age: 10, gender: Female}
	for _, applyOpt := range opts {
		applyOpt(&p)
	}
	return &p
}

func (p People) String() string {
	return fmt.Sprintf("name:%v,age:%v,gender:%v", p.name, p.age, p.gender)
}

func Welcome(people *People) string {
	return fmt.Sprintf("welcome,%s. Your information is %s", people.name, people)
}

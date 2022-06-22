package object

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestPeople_Newyear(t *testing.T) {

	p := new(People)
	fmt.Println(&p)
	p.age = 10
	p.name = "xiaohong"
	x := p.Newyear()
	fmt.Println(p.age, x == p)
	if p.age != 11 {
		t.Errorf("test new year failed")
	}
}

func TestJson(t *testing.T) {
	p := People{name: "xiaot", age: 18, height: 100, weight: 110}
	fmt.Println(&p)
	x := reflect.ValueOf(p)
	fmt.Println(&x)

	a := 10
	fmt.Println(a, reflect.ValueOf(a))
}

func TestNet(t *testing.T) {
	addr, err := net.ResolveIPAddr("ip", "google.cn")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(addr, addr.IP.String())
	}

	lookupAddr, err := net.LookupAddr("8.8.8.8")
	fmt.Println(lookupAddr, err)
}

func TestReflect(t *testing.T) {
	var x float64 = 5.0
	x_t := reflect.TypeOf(x)
	x_v := reflect.ValueOf(x)
	fmt.Println(x_t, x_t.Name(), x_t.String(), x_v,x_v.Type(), x_v.String(), x_v.Kind(), x_v.Kind() == reflect.Float64)
}

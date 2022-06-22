package object

import (
	"fmt"
	"testing"
)

func TestInteger_Add(t *testing.T) {
	var a Integer = 5
	a.Add(10)
	if a != 15 {
		fmt.Println(a)
		t.Errorf("integer add error")
	}
}

func TestSlice(t *testing.T) {
	c := [5]int{1, 2, 3, 4, 5}
	a := c
	b := c[:3]
	fmt.Println(a, b)
	b[0] = 2
	fmt.Println(a, b)

}

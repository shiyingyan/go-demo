package extentions

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	a := [3]string{"1", "2", "3"}
	//b := types.NewArray(types.Int, 10)
	fmt.Println(reflect.TypeOf(a).Kind())
	if v, ok := interface{}(a).([3]string); ok {
		fmt.Println("[]string", v)
	}
	switch interface{}(a).(type) {
	case [5]int:
		fmt.Println("array")
	default:
		fmt.Println("default", reflect.TypeOf(a))
	}

}



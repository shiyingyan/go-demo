package extentions

import (
	"reflect"
)

func Contains(a interface{}, ele interface{}) bool {
	cT := reflect.TypeOf(a).Name()
	eT := reflect.TypeOf(ele).Name()
	if cT != eT {
		return false
	}
	return true
}


package greetings

import (
	"fmt"
	"testing"
)

func TestWelcome(t *testing.T) {
	fmt.Println(Welcome(New("xiaohong")))
	fmt.Println(Welcome(New("xiaoqiang", WithAge(20), WithGender(Male))))

}

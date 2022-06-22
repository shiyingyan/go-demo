package morestrings

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"hello world", "dlrow olleh"},
		{"welcome", "emoclew"},
		{"中国", "国中"},
	}

	for _, tc := range testcases {
		if Reverse(tc.in) != tc.want {
			t.Errorf("%v failed", tc)
		}
	}
}

// need go version >= 1.18
func FuzzReverse(f *testing.F) {
	//testcases := []string{"hello world", "welcome", "中国"}
	testcases := []string{"Hello, world", " ", "!12345", "中国"}
	for _, tc := range testcases {
		//fmt.Printf("seed:%v", tc)
		//f.Logf("seed:%v,rev:%v,doublerev:%v", tc, Reverse(tc), Reverse(Reverse(tc)))
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doublerev := Reverse(rev)
		t.Logf("ori:%v,rev:%v,doublerev:%v", orig, rev, doublerev)
		//fmt.Printf("ori:%v,rev:%v,doublerev:%v", orig, rev, doublerev)
		if orig != doublerev {
			t.Errorf("Before: %q, after: %q", orig, doublerev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(doublerev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}

	})
}

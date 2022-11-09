package hello

import "testing"

func helloTest(t testing.T) {
	want := "hello World!"
	if str := Hello(); want != str {
		t.Errorf("hello():%v ,want: %v\n", str, want)
	}
}

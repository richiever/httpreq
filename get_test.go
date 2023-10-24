package http_req

import (
	"bytes"
	"testing"
)

type Requests1 struct {
	arg1 []byte
	arg2 Requests
}

var R = Requests{"https://jsonplaceholder.typicode.com/posts/1"}
var WantVal []byte = R.Get()
var Allr = []Requests1{
	Requests1{arg1: R.Get(), arg2: R},
}

func Test_Get(t *testing.T) {

	for _, tests := range Allr {
		got := tests.arg2.Get()
		want := tests.arg1
		if !bytes.Equal(got, want) {
			t.Errorf("wanted %s, got %s", want, got)
		}
	}

}

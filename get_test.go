package http_req

import (
	"bytes"
	"testing"
)

var WantVal []byte = R.Get()

var R = Requests{"https://jsonplaceholder.typicode.com/posts/1"}

func Test_Get(t *testing.T) {
	got := R.Get()
	want := WantVal
	if !bytes.Equal(got, want) {
		t.Errorf("wanted %s, got %s", want, got)
	}
}

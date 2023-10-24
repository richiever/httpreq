package http_req

import "testing"

var R = Requests{"https://jsonplaceholder.typicode.com/posts/1"}

func Test_Get(t *testing.T) {
	got := Get(R)
	want := Get(R)
	if got != want {
		t.Errorf("wanted %s, got %s", want, got)
	}
}

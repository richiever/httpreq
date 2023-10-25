package http_req

import (
	"net/http"
	"testing"
)

type Requests2 struct {
	arg1 http.Response
	arg2 RequestsPost
}

var Rtest = RequestsPost{url: "https://jsonplaceholder.typicode.com/posts/1", postBody: map[string]string{"name": "Toby", "email": "Toby@example.com"}, requestType: "application/json"}

var Allr2 = []Requests2{
	Requests2{arg1: Rtest.PostRequestMakeRequest(Rtest.PostRequestByteBuffer()), arg2: Rtest},
}

func Test_Post(t *testing.T) {

	for _, tests := range Allr2 {
		got := Rtest.PostRequestMakeRequest(Rtest.PostRequestByteBuffer())
		want := tests.arg1
		if got.StatusCode != want.StatusCode {
			t.Errorf("not valid")
		}
	}

}

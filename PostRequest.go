package http_req

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestsPost struct {
	url         string
	postBody    map[string]string
	requestType string
}

func (r RequestsPost) PostRequestByteBuffer() bytes.Buffer {
	postB, _ := json.Marshal(r.postBody)

	responseBody := bytes.NewBuffer(postB)
	return *responseBody
}

func (r RequestsPost) PostRequestMakeRequest(b1 bytes.Buffer) http.Response {
	resp, err := http.Post(r.url, r.requestType, &b1)
	if err != nil {
		log.Fatalf("An Error Occured %v, closing", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	return *resp
}

func (r RequestsPost) Post() http.Response {
	return r.PostRequestMakeRequest(r.PostRequestByteBuffer())
}

func (r RequestsPost) PostAsString() string {
	body, err := ioutil.ReadAll(r.Post().Body)
	if err != nil {
		log.Fatalln(err)
	}
	strBody := string(body)
	return strBody

}

package http_req

import (
	"bytes"
	"encoding/json"
	"io"
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

package http_req

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Requests struct {
	url string
}

func (r Requests) RespGet() *http.Response {
	response, err := http.Get(r.url)

	if err != nil {
		log.Fatalln(err)
		return response
	}
	return response
}

func (r Requests) GetBody(response *http.Response) []byte {
	body_req, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body_req
}

func (r Requests) GetStringify(body_req []byte) string {

	stringifiedBody := string(body_req)
	return stringifiedBody
}

func (r Requests) Get() []byte {
	Resp := r.RespGet()
	return r.GetBody(Resp)

}

package http_req

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Requests struct {
	url string
}

func Get(r Requests) string {
	response, err := http.Get(r.url)

	if err != nil {
		log.Fatalln(err)
		return "BAD REQUEST"
	}
	body_req, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
		return "BAD REQUEST"
	}
	stringifiedBody := string(body_req)
	return stringifiedBody
}

package client

import (
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	BaseUri string
}

func (c *HttpClient) Get(path string) []byte {
	client := http.Client{Timeout: time.Second * 10}
	res, err := client.Get(c.BaseUri + path)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return data
}

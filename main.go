package main

import (
	"encoding/json"
	"github.com/limingxinleo/bing-img-download/bing"
	"github.com/limingxinleo/bing-img-download/client"
	"io/ioutil"
	"time"
)

func main() {
	cli := &client.HttpClient{
		BaseUri: "https://cn.bing.com",
	}

	data := cli.Get("/HPImageArchive.aspx?format=js&idx=0&n=1")

	result := &bing.ResultDTO{}
	json.Unmarshal(data, result)

	for _, image := range result.Images {
		bt := cli.Get(image.Url)
		filename := time.Now().Format("2006-01-02") + "." + image.ToMd5() + ".jpg"
		err := ioutil.WriteFile(filename, bt, 0644)
		if err != nil {
			panic(err)
		}
	}
}

package bing

import (
	"crypto/md5"
	"fmt"
)

type ImageDTO struct {
	Url string `json:"url"`
}

type ResultDTO struct {
	Images []ImageDTO `json:"images"`
}

func (i *ImageDTO) ToMd5() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(i.Url)))
}

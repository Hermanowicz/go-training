package models

import "time"

type Link struct {
	OriginalUrl string
	ShortId     string
	Views       int
	Stamp       int64
	Published   bool
}

func NewLink(url string) (*Link, error) {
	shortId := "new short id"
	timeStamp := time.Now().Unix()

	return &Link{
		OriginalUrl: url,
		ShortId:     shortId,
		Views:       0,
		Stamp:       int64(timeStamp),
		Published:   true,
	}, nil
}

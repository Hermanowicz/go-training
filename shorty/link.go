package main

import "time"

type Link struct {
	OrginalUrl string
	ShortId    string
	Views      int
	Stamp      int64
}

func NewLink(url string) (*Link, error) {
	shortId := "new short id"
	timeStamp := time.Now().Unix()

	return &Link{
		OrginalUrl: url,
		ShortId:    shortId,
		Views:      0,
		Stamp:      int64(timeStamp),
	}, nil
}

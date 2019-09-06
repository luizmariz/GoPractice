package main

import (
	"io/ioutil"
)

type page struct {
	Title string
	Body []byte
}

func (p *page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &page{Title: title, Body: body}, nil
}
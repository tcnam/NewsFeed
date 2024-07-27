package main

import (
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	var filename string = p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	var filename string = title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

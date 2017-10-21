package main

import (
	"io/ioutil"
	"strings"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := "pages/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "pages/" + title + ".txt"
	body, err:= ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func (p *Page) PrittyTitle() string {
	return strings.Replace(p.Title, "_", " ", -1)
}
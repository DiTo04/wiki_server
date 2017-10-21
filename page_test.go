package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const testPageName = "testing_page"
const testBody = "Testitest"

func removeFile(s string) {
	os.Remove(s)
}

func createFile(path string) {
	err := ioutil.WriteFile(path, []byte(testBody), 0600)
	if err != nil {
		println("Got error: " + err.Error())
	}
}

func TestPageSave(t *testing.T) {
	//Given
	expectedPath := "pages/" + testPageName + ".txt"
	page := &Page{testPageName, []byte("body")}
	//when
	page.save()
	//Then
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Error("Save did not save the file to: " + expectedPath)
	}
	removeFile(expectedPath)
}

func TestPageLoad(t *testing.T) {
	//Given
	createFile("pages/" + testPageName + ".txt")
	//When
	p, err := loadPage(testPageName)
	switch {
	case err != nil:
		t.Error("Got error while loading: " + err.Error())
	case p.Title != testPageName:
		t.Error("Page title did not match. Expected: " +
			testPageName + ", got: " + p.Title)
	case string(p.Body) != testBody:
		t.Error("Page body do not match. Expected: " +
			testBody + ", got: " + string(p.Body))
	}
}

func TestPageLoadError(t *testing.T) {
	//Given
	file := testPageName + "asd"
	//When
	_, err := loadPage(file)
	if err == nil {
		t.Error("Did not throw error when loading non existent file!")
	}
}

func TestPagePrittyTitle(t *testing.T) {
	//given
	p := &Page{testPageName, []byte(testBody)}
	//When
	s := p.PrittyTitle()
	//Then
	if s != "testing page" {
		t.Error("Expected: testing page, Got: " + s)
	}
}

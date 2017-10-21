package main

import (
	"github.com/DiTo04/wiki_server/mock_http"
	"testing"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/url"
	"fmt"
)

func TestServerMakeHandler(t *testing.T)  {
	controller := gomock.NewController(t)
	defer controller.Finish()
	//Given
	title := "asd"
	mockRW := mock_http.NewMockResponseWriter(controller)
	aUrl := &url.URL{Path:"/view/" + title}
	req := &http.Request{URL:aUrl}
	//When
	var gottenWriter http.ResponseWriter
	var gottenRequest *http.Request
	var gottenString string
	handler := makeHandler(func(writer http.ResponseWriter, r *http.Request, s string) {
		gottenRequest = r
		gottenString = s
		gottenWriter = writer
	})
	handler(mockRW, req)
	//Then
	switch {
	case gottenWriter != mockRW:
		t.Error("Handler did not pass correct writer")
	case gottenRequest != req:
		t.Error("Handler did not pass correct Req")
	case gottenString != title:
		t.Error("Expected: " + title + ", Got: " + gottenString)
	}
}
func TestServerMakeHandlerWrongURL(t *testing.T)  {
	controller := gomock.NewController(t)
	defer controller.Finish()
	//Given
	title := "asd"
	mockRW := mock_http.NewMockResponseWriter(controller)
	aUrl := &url.URL{Path:"/WRONG/" + title}
	req := &http.Request{URL:aUrl}
	//When
	var gottenWriter http.ResponseWriter
	var gottenRequest *http.Request
	var gottenString string
	handler := makeHandler(func(writer http.ResponseWriter, r *http.Request, s string) {
		gottenRequest = r
		gottenString = s
		gottenWriter = writer
	})
	mockRW.EXPECT().Header().Return(http.Header{}).AnyTimes()
	mockRW.EXPECT().WriteHeader(404)
	mockRW.EXPECT().Write(gomock.Any())
	handler(mockRW, req)
	// Then
	didAcceptURL := false
	didAcceptURL = didAcceptURL || (gottenWriter != nil)
	didAcceptURL = didAcceptURL || (gottenRequest != nil)
	didAcceptURL = didAcceptURL || (gottenString != "")
	if (didAcceptURL) {
		t.Error("Handler did continue with wrong url!")
	}
}

func TestRederectToHomePage(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	//Given
	homePageLocation := "/view/Home_Page"
	aUrl := &url.URL{Path:"/"}
	req := &http.Request{URL:aUrl}
	mockRW := mock_http.NewMockResponseWriter(controller)
	header := http.Header{}
	mockRW.EXPECT().Header().Return(header)
	mockRW.EXPECT().WriteHeader(http.StatusFound)
	//When
	redirectToHomePage(mockRW, req)
	redirectLocation := header.Get("Location")
	//Then
	if redirectLocation != homePageLocation {
		t.Error(fmt.Sprintf("Expected %s, but Got: %s",homePageLocation, redirectLocation ))
	}
}
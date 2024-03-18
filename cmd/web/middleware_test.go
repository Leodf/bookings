package main

import (
	"net/http"
	"testing"
)

type myHandlerMock struct{}

func (mh *myHandlerMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func TestNoSurf(t *testing.T) {
	var myH myHandlerMock
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}
func TestSessionLoad(t *testing.T) {
	var myH myHandlerMock
	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}

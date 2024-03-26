package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFormRequire(t *testing.T) {
	r := httptest.NewRequest("POST", "/any_route", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}
	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/any_route", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestFormHas(t *testing.T) {
	r := httptest.NewRequest("POST", "/any_route", nil)
	form := New(r.PostForm)
	hasField := form.Has("a")
	if hasField {
		t.Error("forms return true a field there is no exist")
	}
	postedData := url.Values{}
	postedData.Add("a", "b")

	r, _ = http.NewRequest("POST", "/any_route", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	hasField = form.Has("a")
	if !hasField {
		t.Error("forms return false when should return true to a exist field")
	}
}

func TestFormMinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/any_route", nil)
	form := New(r.PostForm)
	form.MinLength("a", 10)
	if form.Valid() {
		t.Error("form shows valid when minLength fields are missing")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	r, _ = http.NewRequest("POST", "/any_route", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.MinLength("a", 10)
	if form.Valid() {
		t.Error("form shows valid when minLength is not satisfied")
	}

	postedData = url.Values{}
	postedData.Add("a", "aaaaaaaaaa")
	r, _ = http.NewRequest("POST", "/any_route", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.MinLength("a", 10)
	if !form.Valid() {
		t.Error("form shows invalid when minLength is satisfied")
	}

	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestFormIsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/any_route", nil)
	form := New(r.PostForm)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows valid when email field is missing")
	}

	postedData := url.Values{}
	postedData.Add("email", "wrong_email@domain")
	r, _ = http.NewRequest("POST", "/any_route", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows valid when isEmail is not satisfied")
	}

	postedData = url.Values{}
	postedData.Add("email", "wrong_email@domain.com")
	r, _ = http.NewRequest("POST", "/any_route", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows invalid when isEmail is satisfied")
	}
}

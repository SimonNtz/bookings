package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.Form)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}
func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.Form)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.Form = postedData
	form = New(r.Form)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("shows does not have required when it does")
	}
}
func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.Form)

	has := form.Has("whatever", r)
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	r.Form = postedData
	form = New(r.Form)

	has = form.Has("a", r)
	if !has {
		t.Error("shows does not have field when it should")
	}
}

// func TestForm_MinLength(t *testing.T) {
// 	r := httptest.NewRequest("POST", "/whatever", nil)
// 	form := New(r.Form)
// 	postedData := url.Values{}
// 	postedData.Add("a", "abc")

// 	hasMinLength := form.MinLength("a", 4)
// 	if hasMinLength {
// 		t.Error("form shows field has minimum length when it should have been has not")
// 	}

// 	hasMinLength = form.MinLength("a", 3)
// 	if !hasMinLength {
// 		t.Error("shows does have not minimum length when it does")
// 	}
// }

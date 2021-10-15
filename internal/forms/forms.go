package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct and a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		Values: data,
		Errors: errors(map[string][]string{}),
	}
}

// Valid check if the forms has any error
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required check if the form has the given field
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if a field exists in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	if r.Form.Get(field) == "" {
		return false
	}
	return true
}

// MinLength check if the field has the given minimimum length
func (f *Form) MinLength(field string, length int) bool {
	if len(f.Get(field)) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

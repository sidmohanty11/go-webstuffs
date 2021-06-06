package forms

import (
	"net/url"
)

//creates a custom form struct, embeds a url.Values obj
type Form struct {
	url.Values
	Errors errors
}

//initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

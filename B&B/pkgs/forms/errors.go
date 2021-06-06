package forms

import "net/http"

type errors map[string][]string

//adds an err msg for a given form field (first_name,email etc)
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

//returns the first error message
func (e errors) Get(field string) string {
	es := e[field]

	if len(es) == 0 {
		return ""
	}

	return es[0]
}

//checks if it "has" fields in post req
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		return false
	}

	return true
}

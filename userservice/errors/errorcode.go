package errors

// Details struct contains the application logic specific
// error code, readable message and a link to a support doc.
type Details struct {
	Code        string      `json:"code"`
	Description string      `json:"description,omitempty"`
	Link        string      `json:"link,omitempty"`
	Errors      interface{} `json:"errors,omitempty"`
}

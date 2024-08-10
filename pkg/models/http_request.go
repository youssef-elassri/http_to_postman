package models

type Header struct {
	Key   string
	Value string
}
type Request struct {
	Description string
	Method      string
	Url         string
	Headers     []Header
	Body        string
}

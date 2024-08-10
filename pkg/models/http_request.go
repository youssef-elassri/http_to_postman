package models

type Header struct {
	key   string
	value string
}
type Request struct {
	description string
	method      string
	url         string
	headers     []Header
	body        string
}

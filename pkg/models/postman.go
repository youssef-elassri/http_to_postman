package models

type PostmanCollection struct {
	Info CollectionInfo `json:"info"`
	Item []PostmanItem  `json:"item"`
}

type CollectionInfo struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type PostmanItem struct {
	Name    string         `json:"name"`
	Request PostmanRequest `json:"request"`
}

type PostmanRequest struct {
	Method string          `json:"method"`
	Header []PostmanHeader `json:"header"`
	Body   PostmanBody     `json:"body,omitempty"`
	Url    PostmanURL      `json:"url"`
}

type PostmanHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PostmanBodyOptionsRaw struct {
	Language string `json:"language"`
}

type PostmanBodyOptions struct {
	Raw PostmanBodyOptionsRaw `json:"raw"`
}

type PostmanBody struct {
	Mode    string             `json:"mode"`
	Raw     string             `json:"raw"`
	Options PostmanBodyOptions `json:"options"`
}

type PostmanURL struct {
	Raw  string   `json:"raw"`
	Host []string `json:"host"`
	Path []string `json:"path"`
}

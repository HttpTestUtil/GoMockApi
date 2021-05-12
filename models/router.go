package models

const (
	METHOD_GET  = "get"
	METHOD_POST = "post"
)

type Router struct {
	Route    string                 `json:"route"`
	Method   string                 `json:"method"`
	Response map[string]interface{} `json:"response"`
}

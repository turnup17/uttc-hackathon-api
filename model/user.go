package model

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type UserResForHTTPPost struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type UserResForID struct {
	Id string `json:"id"`
}

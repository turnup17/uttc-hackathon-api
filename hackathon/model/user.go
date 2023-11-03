package model

import "time"

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

type KnowledgeResForHTTPGET struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	Date       time.Time `json:"date"`
	Category   int       `json:"category"`
	Details    string    `json:"details"`
	Curriculum int       `json:"curriculum"`
}
type KnowledgeResForHTTPPost struct {
	Name       string `json:"name"`
	Url        string `json:"url"`
	Category   int    `json:"category"`
	Details    string `json:"details"`
	Curriculum int    `json:"curriculum"`
}
type KnowledgeResForID struct {
	Id string `json:"id"`
}
type KnowledgeReqForDelete struct {
	Id string `json:"id"`
}
type KnowledgeReqForHTTPPUT struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	Date       time.Time `json:"date"`
	Category   int       `json:"category"`
	Details    string    `json:"details"`
	Curriculum int       `json:"curriculum"`
}

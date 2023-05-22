package models

type Article struct {
	ID    string `json:id`
	Title string `json:title`
	Body  string `json:body`
}

type User struct {
	userId string `json:userId`
}

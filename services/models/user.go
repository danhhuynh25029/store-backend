package models

type User struct {
	Email    string `json:"email" bson:"email" `
	Name     string `json:"name" bson:"name" `
	Password string `json:"password" bson:"password"`
}

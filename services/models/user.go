package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `id:"id" bson:"_id"`
	Email    string             `json:"email" bson:"email" binding:"required,email"`
	Name     string             `json:"name" bson:"name" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
}

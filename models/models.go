package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `json:"name"		bson:"name"`
	Status string             `json:"status"	bson:"status"`
	UserID string             `json:"user_id"	bson:"user_id"`
}

type User struct {
	ID     primitive.ObjectID 	`bson:"_id"`
	Name   *string             	`json:"username"	bson:"username"`
	Email  *string             	`json:"email"		bson:"email"`
	Password *string             `json:"password"	bson:"password"`
}


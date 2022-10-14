package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"FirstName" json:"FirstName"`
	LastName  string             `bson:"LastName" json:"LastName"`
	Email     string             `bson:"Email" json:"Email"`
	Password  string             `bson:"Password" json:"Password"`
}

func (u *User) HashPassword() {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	u.Password = string(hashed)
}

func (u User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

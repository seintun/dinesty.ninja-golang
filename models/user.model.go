package models

import "github.com/globalsign/mgo/bson"

// User struct represents a basic information of user
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Phone    string        `json:"phone"`
	Orders   []string      `json:"orders"`
}

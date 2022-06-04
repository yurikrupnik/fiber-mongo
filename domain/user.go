package domain

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name,omitempty" validate:"required,min=3,max=36"`
	Role  string `json:"role" bson:"role,omitempty"`
	Age   int    `json:"age" bson:"age,omitempty" validate:"required,numeric"`
	Email string `json:"email" bson:"email,omitempty" validate:"required,email"`
}

type UserSvc interface {
	Get(id string) (*User, error)
	List(category string) ([]*User, error)
	Create(u *User) (*mongo.InsertOneResult, error)
	Delete(id string) error
}

type UserDB interface {
	Get(id string) (*User, error)
	List(category string) ([]*User, error)
	Create(u *User) (*mongo.InsertOneResult, error)
	Delete(id string) error
}

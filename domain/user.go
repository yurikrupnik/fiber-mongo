package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name,omitempty" validate:"required,min=3,max=36"`
	Role  string             `json:"role" bson:"role,omitempty"`
	Age   int                `json:"age" bson:"age,omitempty" validate:"required,numeric" query:"age"`
	Email string             `json:"email" bson:"email,omitempty" validate:"required,email" query:"email"`
}

type UserSvc interface {
	Get(id primitive.ObjectID) (*User, error)
	List(user *User) ([]*User, error)
	Create(u *User) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) error
	Update(id primitive.ObjectID, u *User) error
}

type UserDB interface {
	Get(id primitive.ObjectID) (*User, error)
	List(user *User) ([]*User, error)
	Create(u *User) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) error
	Update(id primitive.ObjectID, u *User) error
}

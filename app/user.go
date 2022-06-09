package app

import (
	"fiber-mongo/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userSvc struct {
	DB domain.UserDB
}

func NewUserSvc(db domain.UserDB) domain.UserSvc {
	return userSvc{
		DB: db,
	}
}

func (us userSvc) Get(id primitive.ObjectID) (*domain.User, error) {
	return us.DB.Get(id)
}

func (us userSvc) List(u *domain.User) ([]*domain.User, error) {
	return us.DB.List(u)
}

func (us userSvc) Create(user *domain.User) (*mongo.InsertOneResult, error) {
	return us.DB.Create(user)
}

func (us userSvc) Delete(id primitive.ObjectID) error {
	return us.DB.Delete(id)
}

func (us userSvc) Update(id primitive.ObjectID, update *domain.User) error {
	return us.DB.Update(id, update)
}

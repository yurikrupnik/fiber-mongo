package app

import (
	"fiber-mongo/domain"
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

func (us userSvc) Get(id string) (*domain.User, error) {
	return us.DB.Get(id)
}

func (us userSvc) List(category string) ([]*domain.User, error) {
	return us.DB.List(category)
}

func (us userSvc) Create(user *domain.User) (*mongo.InsertOneResult, error) {
	return us.DB.Create(user)
}

func (us userSvc) Delete(id string) error {
	return us.DB.Delete(id)
}

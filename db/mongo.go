package db

import (
	"context"
	"fiber-mongo/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mongoStore struct {
	Client *mongo.Client
}

func NewMongoStore() (domain.UserDB, error) {
	//uri := os.Getenv("mongo_uri")
	uri := "mongodb+srv://yurikrupnik:T4eXKj1RBI4VnszC@cluster0.rdmew.mongodb.net/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return mongoStore{Client: client}, nil
}

func (ms mongoStore) Get(id string) (*domain.User, error) {
	col := ms.Client.Database("users").Collection("users")
	filter := bson.M{"id": id}
	user := domain.User{}
	err := col.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ms mongoStore) List(category string) ([]*domain.User, error) {
	col := ms.Client.Database("users").Collection("users")
	filter := bson.M{"category": category}
	log.Println("filter", filter)
	users := []*domain.User{}
	cursor, err := col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (ms mongoStore) Delete(id string) error {
	col := ms.Client.Database("users").Collection("users")
	filter := bson.M{"id": id}
	result, err := col.DeleteOne(context.TODO(), filter)
	log.Println("result", result.DeletedCount)
	return err
}

func (ms mongoStore) Create(user *domain.User) (*mongo.InsertOneResult, error) {
	col := ms.Client.Database("users").Collection("users")
	result, err := col.InsertOne(context.TODO(), user)
	return result, err
}

package db

import (
	"context"
	"fiber-mongo/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	//uri := "mongodb://db/profiless"
	//uri := "http://host.docker.internal:27017"
	//uri := "mongodb://localhost/db" // compose local gow run main.go
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("error connecting mongo")
		return nil, err
	}
	return mongoStore{Client: client}, nil
}

func (ms mongoStore) Get(id primitive.ObjectID) (*domain.User, error) {
	col := ms.Client.Database("users").Collection("users")
	filter := bson.M{"_id": id}
	user := domain.User{}
	err := col.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ms mongoStore) List(u *domain.User) ([]*domain.User, error) {
	col := ms.Client.Database("users").Collection("users")
	// todo generic filter!!!
	log.Println("u", u)
	filter := bson.M{}
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

func (ms mongoStore) Delete(id primitive.ObjectID) error {
	col := ms.Client.Database("users").Collection("users")
	filter := bson.M{"_id": id}
	_, err := col.DeleteOne(context.TODO(), filter)
	return err
}

func (ms mongoStore) Create(user *domain.User) (*mongo.InsertOneResult, error) {
	col := ms.Client.Database("users").Collection("users")
	result, err := col.InsertOne(context.TODO(), user)
	return result, err
}

func (ms mongoStore) Update(id primitive.ObjectID, user *domain.User) error {
	col := ms.Client.Database("users").Collection("users")
	//update := bson.M{"name": user.Name, "email": user.Email, "role": user.Role}
	//fmt.Println(update)
	_, err := col.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": user})
	//log.Println("UpsertedID", result.UpsertedID)
	//_, err := col.UpdateOne(context.TODO(), user)
	return err
}

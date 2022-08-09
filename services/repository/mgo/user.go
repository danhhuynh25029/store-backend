package mgo

import (
	"context"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"store/services/models"
)

type UserRepository interface {
	GetUser(objId string) (*models.User, error)
	AddUser(models.User) error
	UpdateUser(objId string, user models.User) error
	DeleteUser(objId string) error
}

type userRepository struct {
	context    context.Context
	collection *mongo.Collection
	redis      *redis.Client
}

func NewUserRepository(context context.Context, collection *mongo.Collection, redis *redis.Client) UserRepository {
	return &userRepository{context, collection, redis}
}

func (u *userRepository) GetUser(objId string) (*models.User, error) {
	id, _ := primitive.ObjectIDFromHex(objId)
	user := models.User{}
	err := u.collection.FindOne(u.context, bson.M{
		"_id": id,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) AddUser(user models.User) error {
	_, err := u.collection.InsertOne(u.context, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(objId string) error {
	id, _ := primitive.ObjectIDFromHex(objId)
	_, err := u.collection.DeleteOne(u.context, bson.M{
		"_id": id,
	})

	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateUser(objId string, user models.User) error {
	id, _ := primitive.ObjectIDFromHex(objId)
	filter := bson.M{"_id": id}
	update := bson.M{}
	if user.Name != "" {
		update["name"] = user.Name
	}
	if user.Password != "" {
		update["password"] = user.Password
	}
	if user.Email != "" {
		update["email"] = user.Email
	}
	_, _ = u.collection.UpdateOne(
		u.context,
		filter,
		bson.M{
			"$set": update,
		},
	)
	return nil
}

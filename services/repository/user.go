package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"store/services/models"
)

type UserRepository interface {
	GetUser(primitive.ObjectID) (*models.User, error)
	AddUser(models.User) error
	//UpdateUser()
	DeleteUser(id primitive.ObjectID) error
}

type userRepository struct {
	context    context.Context
	collection *mongo.Collection
	redis      *redis.Client
}

func NewUserRepository(context context.Context, collection *mongo.Collection, redis *redis.Client) UserRepository {
	return &userRepository{context, collection, redis}
}

func (u *userRepository) GetUser(id primitive.ObjectID) (*models.User, error) {
	val, err := u.redis.Get(context.Background(), id.Hex()).Result()
	if err != nil {
		fmt.Println("Data is not exists Redis")
		var us models.User
		err := u.collection.FindOne(u.context, bson.M{
			"_id": id,
		}).Decode(&us)
		if err != nil {
			return nil, err
		}
		err = u.redis.Set(context.Background(), us.ID.Hex(), us, 0).Err()
		return &us, nil
	} else {
		var user models.User
		//fmt.Println(val)
		err := json.Unmarshal([]byte(val), &user)
		if err != nil {
			//fmt.Println("Hello world")
			return nil, err
		}
		//fmt.Println(user)
		return &user, nil
	}

}

func (u *userRepository) AddUser(user models.User) error {
	res, err := u.collection.InsertOne(u.context, bson.D{
		{"email", user.Email},
		{"password", user.Password},
		{"name", user.Name},
	})

	if err != nil {
		return err
	}
	fmt.Println(res.InsertedID)
	user.ID, err = primitive.ObjectIDFromHex(res.InsertedID.(primitive.ObjectID).Hex())
	value, _ := json.Marshal(user)
	err = u.redis.Set(context.Background(), res.InsertedID.(primitive.ObjectID).Hex(), value, 0).Err()
	if err != nil {
		fmt.Println("cannot set data into redis")
	}
	return nil
}

func (u *userRepository) DeleteUser(id primitive.ObjectID) error {

	res, err := u.collection.DeleteOne(u.context, bson.M{
		"_id": id,
	})

	if err != nil {
		return err
	}
	fmt.Println(res.DeletedCount)
	return nil
}

//func (u *userRepository) UpdateUser(user models.User) error {
//	res, err := u.collection.UpdateOne(u.context, bson.M{
//		"_id": user.ID,
//	},
//		bson.D{
//			{"$set",bson.D{
//
//			}},
//		}
//	)
//}

package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"store/pkg/config"
	"store/pkg/redis"
	"store/services/domain/user/delivery/http"
	"store/services/domain/user/usecase"
	"store/services/repository/mgo"
)

func main() {

	config, err := config.LoadConfig(".")
	//
	redis := redis.NewRedisClient(config)
	ctx := context.TODO()
	fmt.Println(config.MONGO_LOCAL_URL)
	// Get Collection
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MONGO_LOCAL_URL))
	if err != nil {
		log.Fatal("Cannot connect mongodb")
	}

	userCollection := client.Database("store").Collection("users")

	r := gin.Default()

	group := r.Group("/api/v1")

	userRepo := mgo.NewUserRepository(ctx, userCollection, redis)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHanlder := http.NewUserHandler(userUsecase)
	userRouter := http.NewUserRouter(userHanlder)
	userRouter.UseRoute(group)

	if err != nil {
		log.Fatal("Cannot load file .env")

	}
	if err := r.Run(":" + config.PORT); err != nil {
		panic(err)
	}
}

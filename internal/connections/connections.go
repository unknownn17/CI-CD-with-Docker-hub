package connections

import (
	"conn/internal/api/handler"
	"conn/internal/database/methods"
	"conn/internal/database/service"
	interface17 "conn/internal/interface"
	"context"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDatabase() interface17.TaskService {
	opt, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27018"))
	if err != nil {
		log.Fatal(err)
	}
	if err := opt.Ping(context.Background(), options.Client().ReadPreference); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	data := opt.Database("cicd").Collection("cicddd")
	return &methods.Database{Db: data, Ctx: ctx, Generated: map[int]bool{}, Rng: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func NewService() *service.Service {
	a := NewDatabase()
	return &service.Service{I: a}
}

func NewHandler() *handler.Handler {
	a := NewService()
	return &handler.Handler{S: a}
}

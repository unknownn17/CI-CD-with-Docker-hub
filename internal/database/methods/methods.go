package methods

import (
	"conn/internal/models"
	"context"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Db        *mongo.Collection
	Ctx       context.Context
	Generated map[int]bool
	Rng       *rand.Rand
}

func (u *Database) CreateTask(req *models.TaskCreate) (*models.Task, error) {
	new := models.Task{
		ID:          u.GenerateUniqueRandomNumber(),
		Title:       req.Title,
		Description: req.Description,
		Status:      "created",
		Created_at:   time.Now().Format(time.RFC3339),
	}
	res, err := u.Db.InsertOne(u.Ctx, new)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	insertedid := res.InsertedID

	var res1 models.Task
	if err := u.Db.FindOne(u.Ctx, bson.M{"_id": insertedid}).Decode(&res1); err != nil {
		log.Println(err)
		return nil, err
	}
	return &res1, nil
}

func (u *Database) GetTask(req *models.Task_Get_Delete) (*models.Task, error) {
	var res models.Task
	if err := u.Db.FindOne(u.Ctx, bson.M{"id": req.ID}).Decode(&res); err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}

func (u *Database) GetTasks() ([]*models.Task, error) {
	rows, err := u.Db.Find(u.Ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var res []*models.Task

	for rows.Next(u.Ctx) {
		var all models.Task
		if err := rows.Decode(&all); err != nil {
			log.Println(err)
			return nil, err
		}
		res = append(res, &all)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *Database) UpdateTask(req *models.TaskUpdate) (*models.Task, error) {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.M{"id": req.ID}
	update := bson.M{"$set": bson.M{"title": req.Title, "description": req.Description, "status": "updated"}}

	var updatedDocument models.Task
	err := u.Db.FindOneAndUpdate(u.Ctx, filter, update, opts).Decode(&updatedDocument)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &updatedDocument, nil
}

func (u *Database) DeleteTask(req *models.Task_Get_Delete) error {
	_, err := u.Db.DeleteOne(u.Ctx, bson.M{"id": req.ID})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Database) GenerateUniqueRandomNumber() int {
	if len(u.Generated) >= (10000 - 1 + 1) {
		log.Printf("no more unique numbers available in the range")
	}

	for {
		num := u.Rng.Intn(10000-1+1) + 1
		if !u.Generated[num] {
			u.Generated[num] = true
			return num
		}
	}
}

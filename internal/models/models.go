package models

type TaskCreate struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

type Task struct {
	ID          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"`
	Created_at  string `json:"created_at" bson:"created_at"`
}

type TaskUpdate struct {
	ID          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

type Task_Get_Delete struct {
	ID int `json:"id" bson:"id"`
}

package interface17

import "conn/internal/models"


type TaskService interface{
	CreateTask(req *models.TaskCreate)(*models.Task,error)
	GetTask(req *models.Task_Get_Delete)(*models.Task,error)
	GetTasks()([]*models.Task,error)
	UpdateTask(req*models.TaskUpdate)(*models.Task,error)
	DeleteTask(req *models.Task_Get_Delete)error
}
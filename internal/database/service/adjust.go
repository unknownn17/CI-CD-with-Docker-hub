package service

import (
	interface17 "conn/internal/interface"
	"conn/internal/models"
)

type Service struct {
	I interface17.TaskService
}

func (u *Service) CreateTask(req *models.TaskCreate) (*models.Task, error) {
	return u.I.CreateTask(req)
}

func (u *Service) GetTask(req *models.Task_Get_Delete) (*models.Task, error) {
	return u.I.GetTask(req)
}

func (u *Service) GetTasks() ([]*models.Task, error) {
	return u.I.GetTasks()
}

func (u *Service) UpdateTask(req *models.TaskUpdate) (*models.Task, error) {
	return u.I.UpdateTask(req)
}

func (u *Service) DeleteTask(req *models.Task_Get_Delete) error {
	return u.I.DeleteTask(req)
}

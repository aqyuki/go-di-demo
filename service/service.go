package service

import (
	"context"
	"di/model"
	"di/repository"
)

type (
	TaskService interface {
		CreateTask(ctx context.Context, task *model.Task)
	}

	taskService struct {
		tr repository.TaskRepository
	}
)

func NewTaskService(tr repository.TaskRepository) TaskService {
	return &taskService{
		tr: tr,
	}
}

func (ts *taskService) CreateTask(ctx context.Context, task *model.Task) {
	ts.tr.CreateTask(ctx, task)
}

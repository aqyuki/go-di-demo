package repository

import (
	"context"
	"database/sql"
	"di/model"
)

type (
	TaskRepository interface {
		CreateTask(ctx context.Context, task *model.Task)
	}

	taskRepository struct {
		db *sql.DB
	}
)

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (td *taskRepository) CreateTask(ctx context.Context, task *model.Task) {
	td.db.Exec("INSERT INTO task (title,note) VALUES (?,?);", task.Title, task.Description)
}

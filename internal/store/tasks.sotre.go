package store

import (
	"database/sql"

	"github.com/s3niffer/taskmanagementapp/internal/models"
)

type TasksStore struct {
	db *sql.DB
}

func NewTasksStore(db *sql.DB) TasksStore {
	return TasksStore{
		db: db,
	}
}

func (t TasksStore) CreateTask(task *models.Task, userId int) error {

	query := `
	INSERT INTO tasks (title,description,due_date,user_id)
	VALUES ($1,$2,$3,$4)
	RETURNING id,created_at,status;
	`

	err := t.db.QueryRow(query, task.Title, task.Description, task.Due_date, userId).Scan(&task.ID, &task.Created_at, &task.Status)
	if err != nil {
		return err
	}

	return nil
}

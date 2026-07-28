package store

import (
	"database/sql"
	"fmt"
	"strings"

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

func (t TasksStore) GetTasks(userId int) ([]models.Task, error) {
	var tasks []models.Task
	query := `
		SELECT * FROM tasks
		WHERE user_id = $1;
	`

	rows, err := t.db.Query(query, userId)
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.ID, &task.User_ID, &task.Title, &task.Description, &task.Status, &task.Due_date, &task.Created_at)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return tasks, nil
}

func (t TasksStore) GetTask(taskId int) (models.Task, error) {
	var task models.Task

	query := `
		SELECT * FROM tasks
		WHERE id = $1
	`

	err := t.db.QueryRow(query, taskId).Scan(
		&task.ID,
		&task.User_ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Due_date,
		&task.Created_at,
	)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (t TasksStore) DeleteTask(taskId int) error {
	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	r, err := t.db.Exec(query, taskId)
	if err != nil {
		return err
	}
	num, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return fmt.Errorf("couldn't find any task with id: %d", taskId)
	}

	return nil
}

func (t TasksStore) UpdateTask(task *models.Task) error {
	var (
		sets []string
		args []any
	)

	if task.Description != "" {
		sets = append(sets, fmt.Sprintf("description = $%d", len(sets)+1))
		args = append(args, task.Description)
	}
	if task.Title != "" {
		sets = append(sets, fmt.Sprintf("title = $%d", len(sets)+1))
		args = append(args, task.Title)
	}
	if task.Status != "" {
		sets = append(sets, fmt.Sprintf("status = $%d", len(sets)+1))
		args = append(args, task.Status)
	}
	if task.Due_date != "" {
		sets = append(sets, fmt.Sprintf("due_date = $%d", len(sets)+1))
		args = append(args, task.Due_date)
	}

	if len(sets) != 0 {
		query := fmt.Sprintf(`
		UPDATE tasks
		SET %s
		WHERE id = $%d
		RETURNING *;
		`, strings.Join(sets, ", "),
			len(args)+1,
		)

		args = append(args, task.ID)

		err := t.db.QueryRow(query, args...).Scan(&task.ID, &task.User_ID, &task.Title, &task.Description, &task.Status, &task.Due_date, &task.Created_at)
		if err != nil {
			return err
		}

		return nil
	}

	// no updates needed just get task
	currentTask, err := t.GetTask(task.ID)
	if err != nil {
		return err
	}

	task = &currentTask
	return nil
}

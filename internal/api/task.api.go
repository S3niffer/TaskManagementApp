package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/s3niffer/taskmanagementapp/internal/models"
	"github.com/s3niffer/taskmanagementapp/internal/store"
)

type TasksApi struct {
	store store.TasksStore
}

func NewTasksApi(store store.TasksStore) TasksApi {
	return TasksApi{
		store: store,
	}
}

func (t TasksApi) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if task.Description == "" || task.Title == "" || task.Due_date == "" {
		http.Error(w, "tasks values can't be empty required(title,description,due_date)", http.StatusInternalServerError)
		return
	}

	userId, err := GetUserIdFromRequest(r)
	if err != nil {
		http.Error(w, "could'nt convert user id to int.", http.StatusInternalServerError)
		return
	}

	task.User_ID = userId

	err = t.store.CreateTask(&task, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (t TasksApi) GetTasks(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdFromRequest(r)
	if err != nil {
		http.Error(w, "could'nt convert user id to int.", http.StatusInternalServerError)
		return
	}

	tasks, err := t.store.GetTasks(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

func (t TasksApi) GetTask(w http.ResponseWriter, r *http.Request) {}

func (t TasksApi) DeleteTask(w http.ResponseWriter, r *http.Request) {}

func (t TasksApi) UpdateTask(w http.ResponseWriter, r *http.Request) {}

func GetUserIdFromRequest(r *http.Request) (int, error) {
	return strconv.Atoi(fmt.Sprint(r.Context().Value(models.AuthMiddleUserIdKey{})))
}

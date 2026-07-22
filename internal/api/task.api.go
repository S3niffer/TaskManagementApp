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

	userId, err := strconv.Atoi(fmt.Sprint(r.Context().Value(models.AuthMiddleUserIdKey{})))
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

func DeleteTask(w http.ResponseWriter, r *http.Request) {}

func UpdateTask(w http.ResponseWriter, r *http.Request) {}

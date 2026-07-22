package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (t TasksApi) GetTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	if idStr == "" {
		http.Error(w, "Bad request task id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad request task id", http.StatusBadRequest)
		return
	}

	task, err := t.store.GetTask(id)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, fmt.Sprintf("Not found any task with id %d", id), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(task)
}

func (t TasksApi) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	if idStr == "" {
		http.Error(w, "Bad request task id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad request task id", http.StatusBadRequest)
		return
	}

	err = t.store.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(fmt.Sprintf("task with id %d has been deleted successfully.", id))
}

func (t TasksApi) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.store.UpdateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(task)
}

func GetUserIdFromRequest(r *http.Request) (int, error) {
	return strconv.Atoi(fmt.Sprint(r.Context().Value(models.AuthMiddleUserIdKey{})))
}

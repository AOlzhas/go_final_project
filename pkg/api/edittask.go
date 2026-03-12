package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeError(w, errors.New("не указан идентификатор"))
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, task)
}

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeError(w, err)
		return
	}

	if task.ID == "" {
		writeError(w, errors.New("не указан идентификатор"))
		return
	}

	if task.Title == "" {
		writeError(w, errors.New("не указан заголовок задачи"))
		return
	}

	if err := checkDate(&task); err != nil {
		writeError(w, err)
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, map[string]string{})
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeError(w, errors.New("не указан идентификатор"))
		return
	}

	if err := db.DeleteTask(id); err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, map[string]string{})
}

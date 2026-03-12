package api

import (
	"errors"
	"net/http"
	"time"

	"go_final_project/pkg/db"
)

func doneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			writeError(w, err)
			return
		}
		writeJSON(w, map[string]string{})
		return
	}

	next, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeError(w, err)
		return
	}

	if err := db.UpdateDate(next, id); err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, map[string]string{})
}

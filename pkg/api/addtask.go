package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"go_final_project/pkg/db"
)

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(dateFormat)
	}

	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return err
	}

	if task.Repeat != "" {
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}

		if beforeNow(t, now) {
			task.Date = next
		}
	} else {
		if beforeNow(t, now) {
			task.Date = now.Format(dateFormat)
		}
	}

	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeError(w, err)
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

	id, err := db.AddTask(&task)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, map[string]string{
		"id": strconv.FormatInt(id, 10),
	})
}

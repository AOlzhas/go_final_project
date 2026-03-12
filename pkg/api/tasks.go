package api

import (
	"net/http"

	"go_final_project/pkg/db"
)

type TasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, TasksResponse{
		Tasks: tasks,
	})
}

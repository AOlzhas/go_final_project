package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, err error) {
	writeJSON(w, map[string]string{
		"error": err.Error(),
	})
}

func normalizeDate(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func afterNow(a, b time.Time) bool {
	return normalizeDate(a).After(normalizeDate(b))
}

func beforeNow(a, b time.Time) bool {
	return normalizeDate(a).Before(normalizeDate(b))
}

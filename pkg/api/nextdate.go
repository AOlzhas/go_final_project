package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat не указан")
	}

	date, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", err
	}

	parts := strings.Split(repeat, " ")
	if len(parts) == 0 {
		return "", errors.New("неверный формат repeat")
	}

	switch parts[0] {
	case "d":
		if len(parts) != 2 {
			return "", errors.New("неверный формат repeat")
		}

		days, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", errors.New("неверный формат repeat")
		}
		if days < 1 || days > 400 {
			return "", errors.New("неверный формат repeat")
		}

		date = date.AddDate(0, 0, days)
		for !afterNow(date, now) {
			date = date.AddDate(0, 0, days)
		}

		return date.Format(dateFormat), nil

	case "y":
		if len(parts) != 1 {
			return "", errors.New("неверный формат repeat")
		}

		date = date.AddDate(1, 0, 0)
		for !afterNow(date, now) {
			date = date.AddDate(1, 0, 0)
		}

		return date.Format(dateFormat), nil

	default:
		return "", errors.New("неподдерживаемый формат repeat")
	}
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var now time.Time
	var err error

	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(dateFormat, nowStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	next, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, _ = w.Write([]byte(next))
}

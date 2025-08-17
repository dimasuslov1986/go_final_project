package api

import (
	"encoding/json"
	"errors"
	"go_final_project/pkg/consts"
	"go_final_project/pkg/db"
	"net/http"
	"strconv"
	"time"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "неверный формат JSON"})
		return
	}

	if task.Title == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "не указан заголовок задачи"})
		return
	}

	if err := checkDate(&task); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "Ошибка добавления задачи: " + err.Error()})
		return
	}

	writeJson(w, http.StatusOK, map[string]string{"id": strconv.FormatInt(id, 10)})

}

// checkDate
func checkDate(task *db.Task) error {

	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(consts.DateFormat)
	}

	t, err := time.Parse(consts.DateFormat, task.Date)
	if err != nil {
		return errors.New("неверный формат даты")
	}

	if task.Repeat != "" {

		if task.Date == now.Format(consts.DateFormat) {
			return nil
		}

		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return errors.New("неверное правило повторения")
		}

		if afterNow(now, t) {
			if len(task.Repeat) == 0 {
				// если правила повторения нет, то берём сегодняшнее число
				task.Date = now.Format(consts.DateFormat)
			} else {
				// в противном случае, берём вычисленную ранее следующую дату
				task.Date = next
			}
		}
	} else {
		task.Date = now.Format(consts.DateFormat)
	}
	return nil
}

func writeJson(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "ошибка записи Json", http.StatusInternalServerError)
	}
}

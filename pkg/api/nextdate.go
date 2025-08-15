package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go_final_project/pkg/consts"
)

// NextDate возвращает следующую дату от заданной по правилу repeat
func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	date, err := time.Parse(consts.DateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("ошибка формата даты начала отсчёта повторений: %w", err)
	}

	repRule := strings.Split(repeat, " ")

	switch repRule[0] {
	case "d":

		if len(repRule) != 2 {
			return "", fmt.Errorf("ошибка 1 формата правила повторений в днях: %w", err)
		}
		// получаем интервал повторений в днях
		num, err := strconv.Atoi(repRule[1])
		if err != nil {
			return "", fmt.Errorf("ошибка преобразования : %w", err)
		}

		// проверяем чтобы интервал для повторений в днях не был больше 400
		if num > 400 || num < 1 {
			return "", fmt.Errorf("ошибка 2 формата правила повторений: %w", err)
		}
		for {
			date = date.AddDate(0, 0, num)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(consts.DateFormat), nil

	case "y":

		if len(repRule) != 1 {
			return "", fmt.Errorf("ошибка 1 формата правила повторений в днях: %w", err)
		}
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(consts.DateFormat), nil

	default:
		return "", fmt.Errorf("ошибка 3 формата правила повторений: %w", err)
	}
}

func afterNow(date, now time.Time) bool {

	if date.After(now) {
		return true
	}
	return false
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	nowStr := r.FormValue("now")
	dstart := r.FormValue("date")
	repeat := r.FormValue("repeat")

	if nowStr == "" {
		nowStr = time.Now().Format(consts.DateFormat)
	}

	now, err := time.Parse(consts.DateFormat, nowStr)
	if err != nil {
		fmt.Errorf("ошибка формата даты в запросе клиента: %w", err)
		return
	}

	nextDate, err := NextDate(now, dstart, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(nextDate))

}

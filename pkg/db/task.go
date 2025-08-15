package db

import "fmt"

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// AddTask добавляет задачу в таблицу scheduler и возвращает id добавленной записи
func AddTask(task *Task) (int64, error) {
	var id int64

	query := "INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)"

	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

// Tasks возвращает заданное limit кол-во записей задач
func Tasks(limit int) ([]*Task, error) {

	query := "SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT ?"

	row, err := db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("ошибка  : %w", err)
	}

	defer row.Close()

	var tasks []*Task

	for row.Next() {
		var task Task

		err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("ошибка  : %w", err)
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// GetTask возвращает запись задачи по заданному id
func GetTask(id string) (*Task, error) {

	var task Task

	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?"

	err := db.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// UpdateTask редактирует поля date, title, comment, repeat записи задачи
func UpdateTask(task *Task) error {

	query := "UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?"

	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

// DeleteTask удаляет запись задачи по заданному id
func DeleteTask(id string) error {

	query := "DELETE FROM scheduler WHERE id = ?"

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for deleting task`)
	}
	return nil
}

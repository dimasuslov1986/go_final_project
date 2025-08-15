package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);
`

// Init подключается к базе данных
func Init(dbFile string) error {

	// проверяем существование файла
	_, err := os.Stat(dbFile)
	var install bool
	if err != nil {
		install = true
	}

	// подключаем БД
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("ошибка подключения БД: %w", err)
	}

	// err = db.Ping()
	// if err != nil {
	// 	 log.Fatal(err)
	// }

	// создаем таблицу и индекс если install равен true
	if install {

		_, err = db.Exec(schema)
		if err != nil {
			return fmt.Errorf("ошибка создания схемы БД: %w", err)
		}
	}

	return nil
}

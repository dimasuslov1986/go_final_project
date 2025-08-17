package server

import (
	"fmt"
	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
	"net/http"
)

func Run() error {

	var closer func()

	closer, err := db.Init("scheduler.db")
	if err != nil {
		return err
	}

	defer closer()

	api.Init()

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	fmt.Println(`Запускаем сервер c приложением "Планировщик задач" на порту 7540`)
	err = http.ListenAndServe(":7540", nil)
	if err != nil {
		fmt.Printf("Start server error: %s", err.Error())
	}

	return nil
}

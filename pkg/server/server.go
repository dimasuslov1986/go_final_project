package server

import (
	"fmt"
	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
	"net/http"
)

func Run() error {

	if err := db.Init("scheduler.db"); err != nil {
		return err
	}

	api.Init()

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	fmt.Println("Запускаем сервер")
	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		fmt.Printf("Start server error: %s", err.Error())
	}

	return nil
}

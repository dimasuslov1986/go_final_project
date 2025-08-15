package main

import (
	"fmt"
	"go_final_project/pkg/server"
)

func main() {

	if err := server.Run(); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}

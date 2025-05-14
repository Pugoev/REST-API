package main

import (
	"fmt"
	"net/http"
)

func main() {
	sum = 0 // Инициализация глобального счетчика

	http.HandleFunc("/api/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getNotesHandler(w, r)
		} else if r.Method == http.MethodPost {
			createNoteHandler(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
		return
	}

	fmt.Println("Сервер запущен на порту 8080")
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getNotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var noteSlice []Note
	for _, v := range notes {
		noteSlice = append(noteSlice, *v)
	}
	json.NewEncoder(w).Encode(noteSlice)
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var note Note
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(bodyBytes, &note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	sum++ // Увеличиваем глобальный счетчик
	note.ID = fmt.Sprintf("%d", sum)
	notes[note.ID] = &note // Сохраняем заметку с новым ID

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note) // Отправляем созданную заметку в ответе
}

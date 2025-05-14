package main

var sum int // Глобальный счетчик для ID

type Note struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

var notes = make(map[string]*Note) // Хранилище заметок

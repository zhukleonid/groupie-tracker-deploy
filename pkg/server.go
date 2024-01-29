package pkg

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var Mux *http.ServeMux

func Server() *http.Server {
	// Мультиплексор HTTP-запросов для маршрутизации
	// запросов к соотвествующим обработчикам
	Mux = http.NewServeMux()

	// Регистрация oбработчика домашней страницы
	Mux.HandleFunc("/", HomeHandler)

	// Регистрация обработчика страницы с информацией о группах и артистах
	Mux.HandleFunc("/band", BandHandler)

	fileServer := http.FileServer(http.Dir("web/static"))

	Mux.Handle("/web/static/", http.StripPrefix("/web/static/", fileServer))

	// Хранит конфигурации запуска HTTP-сервера
	S := &http.Server{
		Addr:         ":5000",           // адрес прослушки (хост и порт)
		ReadTimeout:  30 * time.Second,  // время ожидания для чтения данных
		WriteTimeout: 90 * time.Second,  // время ожидания для записи данных
		IdleTimeout:  120 * time.Second, // время простоя
		Handler:      Mux,               // установка обрабочтиков для входящих запросов
	}

	fmt.Printf("Тестовый URL: %s"+"\n", "http://localhost:5000")
	fmt.Printf("Основной URL: %s"+"\n", "http://localhost:8080")
	go func() {
		err := S.ListenAndServe() // Метод структуры Server для запуска сервера
		if err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()
	return S
}

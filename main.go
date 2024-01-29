package main

import (
	"log"
	"lzhuk/groupie-tracker/pkg"
	"net/http"
	"os"
)

func main() {
	
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Не удалось открыть файл лога:", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.SetPrefix("LOG_INFO: ")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	server := pkg.Server()

	server = &http.Server{
		Addr:         ":8080",             // адрес прослушки (хост и порт)
		ReadTimeout:  server.ReadTimeout,  // время ожидания для чтения данных
		WriteTimeout: server.WriteTimeout, // время ожидания для записи данных
		IdleTimeout:  server.IdleTimeout,  // время простоя
		Handler:      pkg.Mux,
	}
	err = server.ListenAndServe() // Метод структуры Server для запуска сервера
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}

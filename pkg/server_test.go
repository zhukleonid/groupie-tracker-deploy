package pkg_test

import (
	"lzhuk/groupie-tracker/pkg"
	"testing"
)

func TestServer(t *testing.T) {
	server := pkg.Server()
	defer server.Close()

	// Проверяем, что сервер запущен
	if server == nil {
		t.Error("Сервер не запущен")
	}
}

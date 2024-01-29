package pkg_test

import (
	"encoding/json"
	"lzhuk/groupie-tracker/pkg"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetBandInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testData := []pkg.Band{
			{
				ID:           1,
				Image:        "non",
				Name:         "Lzhuk",
				Members:      nil,
				CreationDate: 1,
				FirstAlbum:   "Dec",
				Locations:    "Astana",
				ConcertDates: "24-01-2024",
				Relations:    "1",
			},
			{
				ID:           2,
				Image:        "non",
				Name:         "Tsvetash",
				Members:      nil,
				CreationDate: 1,
				FirstAlbum:   "Dec",
				Locations:    "Almata",
				ConcertDates: "25-01-2024",
				Relations:    "2",
			},
		}

		json.NewEncoder(w).Encode(testData)
	}))
	defer server.Close()

	// Тест 1 для проверки корректности получения, обработки и декодирования данных из формата JSON в структуру данных программы
	artistAPI := server.URL

	bands, err := pkg.GetBandInfo(artistAPI)
	if err != nil {
		t.Errorf("Ожидалось не ошибка, однако вывод %v", err)
	}

	expectedData := []pkg.Band{
		{
			ID:           1,
			Image:        "non",
			Name:         "Lzhuk",
			Members:      nil,
			CreationDate: 1,
			FirstAlbum:   "Dec",
			Locations:    "Astana",
			ConcertDates: "24-01-2024",
			Relations:    "1",
		},
		{
			ID:           2,
			Image:        "non",
			Name:         "Tsvetash",
			Members:      nil,
			CreationDate: 1,
			FirstAlbum:   "Dec",
			Locations:    "Almata",
			ConcertDates: "25-01-2024",
			Relations:    "2",
		},
	}

	if !reflect.DeepEqual(bands, expectedData) {
		t.Errorf("Ожидалось %v, вывод %v", expectedData, bands)
	}

	// Тест 2 для проверки возврата ошибки функции при передаче невалидного URL к API
	artistAPI = "http/"

	bands, err = pkg.GetBandInfo(artistAPI)

	if err == nil {
		t.Errorf("Ожидалась ошибка, однако вывод %v", err)
	}

	// Тест 3 для проверки процесса декодирования данных полученных из API
	type BandTestAdress struct {
		Adress       string   `json:"adress"`
	}
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bandTestAdress := BandTestAdress{
				Adress:       "hhhh",
			}
		json.NewEncoder(w).Encode(bandTestAdress)
	}))
	defer server.Close()

	artistAPI = server.URL

	_, err = pkg.GetBandInfo(artistAPI)
	if err == nil {
		t.Errorf("Ожидалась ошибка, однако вывод %v", err)
	}
}

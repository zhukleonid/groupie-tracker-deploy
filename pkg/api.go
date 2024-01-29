package pkg

import (
	"encoding/json"
	"net/http"
)

var bands []Band // Хранит срез со структурой данных о группах

type Band struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func GetBandInfo(artistAPI string) ([]Band, error) {
	response, err := http.Get(artistAPI)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&bands)
	if err != nil {
		return nil, err
	}
	return bands, nil
}

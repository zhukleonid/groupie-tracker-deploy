package pkg

import "fmt"

func GetBandByID(id string, s []Band) (*Band, error) {

	for _, b := range s {
		if fmt.Sprintf("%d", b.ID) == id {
			return &b, nil
		}
	}

	return nil, fmt.Errorf("Группа с идентификатором ID= %s не найдена", id)
}

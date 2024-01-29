package pkg_test

import (
	"testing"

	"lzhuk/groupie-tracker/pkg"
)

func Test_GetBandByID(t *testing.T) {
	band := []pkg.Band{
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

	// Тест 1 для успешного нахождения группы
	result, err := pkg.GetBandByID("2", band)
	if err != nil {
		t.Errorf("Не ожидалась ошибка: %v", err)
	}
	if result == nil {
		t.Error("Ожидалось что группа или артист будут найдены, но поиск по полученным данным  из API не дал результатов")
	} else if result.ID != 2 {
		t.Errorf("Ожидалось найти гуппу с индетификатором ID=2, но результат %d", result.ID)
	}

	// Тест 2 для группы которая отсутсвует в структуре данных
	result, err = pkg.GetBandByID("4", band)
	if err == nil {
		t.Error("Ожидалась ошибка, но ошибки нет")
	}
	if result != nil {
		t.Errorf("Ожидался пустой результат поиска, но фактически %+v", result)
	}
}

package dwb

import "errors"

func ReadRec(id_rec int) (int, error) {
	// Функция читает данные из БД
	// id ReadRec--1
	if id_rec < 0 {
		return 0, errors.New(
			"id ReadRec--1: недопустимый идентификатор записи")
	}
	return id_rec, nil
}

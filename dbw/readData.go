package dwb

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)


// Инициализация структуры для работы с записами
func init(){
	type todo_record struct{
		id,
		data_of_creations,
		date_max,
		todo_text,
		is_gone,
		date_of_gone
	}
}
// ЗАГЛУШКА: Читаю данные из БД
func ReadRec(id_rec int) (int, error) {
	// Функция читает данные из БД
	// id ReadRec--1

	if id_rec < 0 {
		return 0, errors.New(
			"id ReadRec--1: недопустимый идентификатор записи")
	}
	db_connect, err := sql.Open("sqlite3", "ep20231204_todo_cli.db")
	// id ReadRec--2:
	if err != nil {
		fmt.Printf("ReadRec--2: Не могу открыть БД %v", err)
		fmt.Scanln()
		os.Exit(1)
	}
	defer db_connect.Close()
	query := `
		SELECT 
			id,
			data_of_creations,
			date_max,
			todo_text,
			is_gone,
			date_of_gone
		from my_todo_list
	`
	rows, err := db_connect.Query(query)
	// id ReadRec--3
	if err != nil {

		fmt.Printf("ReadRec--3: Не могу прочитать из БД %v", err)
		fmt.Scanln()
		os.Exit(1)
	}
	defer rows.Close()
	
	fmt.Printf("rows type is %T/n", rows)



	return id_rec, nil
}

/*/ package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	// Строка подключения
	connStr := "postgres://postgres:vaznia1234@localhost:5432/Wb?sslmode=disable"

	// Подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Ошибка подключения к БД: %v", err))
	}
	defer db.Close()

	// Проверка доступности базы данных
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("База данных недоступна: %v", err))
	}

	fmt.Println("Подключение к базе данных установлено")
}
/*/
package main

import (
	"fmt"
	"GOLANG_L0/internal/database"
)


func main() {
	// Подключение к базе данных
	db := database.ConnectToDB()

	// Автоматическая миграция таблиц
	database.MigrateTables(db)

	fmt.Println("Таблицы успешно созданы и настроены!")
}

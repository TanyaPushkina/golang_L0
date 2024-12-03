package database

import (
	"log"

	"GOLANG_L0/internal/models" // Импорт моделей
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Подключение к базе данных
func ConnectToDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=vaznia1234 dbname=Wb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	return db
}

// Миграция таблиц
func MigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Payment{}, &models.Item{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
}

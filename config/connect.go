package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/ahmadammarm/go-rest-api/internals/models"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: Gagal memuat file .env")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Gagal terkoneksi dengan database:", err)
		return nil, err
	}

	DB = db

	if err := DB.AutoMigrate(&models.Category{}, &models.Product{}); err != nil {
		log.Println("Peringatan: Gagal melakukan migrasi database:", err)
	}

	fmt.Println("Database berhasil terkoneksi")
	return DB, nil
}

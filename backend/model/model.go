package model

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Goly struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	URL      string `json:"url" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
}

func Setup() {

	// dsn := os.Getenv("DSN")
	dsn := "postgres://admin:test@" + os.Getenv("DB_HOST") + ":5432/admin?sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Goly{})
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MyRow struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&MyRow{})

	// Insert a row
	row := MyRow{Name: "John"}
	db.Create(&row)

	// Find all rows
	var rows []MyRow
	db.Find(&rows)

	for _, r := range rows {
		log.Printf("id=%d, name=%s", r.ID, r.Name)
	}
}

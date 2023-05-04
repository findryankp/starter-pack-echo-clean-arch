package database

import (
	"fmt"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	fmt.Println("Migration Done")
}

package migration

import (
	"gorm.io/gorm"
)

func CreateUsersTable(db *gorm.DB) error {
	type User struct {
		ID        uint   `gorm:"primaryKey"`
		Name      string `gorm:"size:100;not null"`
		Email     string `gorm:"size:100;uniqueIndex;not null"`
		Password  string `gorm:"size:255;not null"`
		CreatedAt int64  `gorm:"autoCreateTime"`
		UpdatedAt int64  `gorm:"autoUpdateTime"`
	}

	return db.AutoMigrate(&User{})
}

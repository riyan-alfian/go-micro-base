package seed

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Users(db *gorm.DB) error {
	users := []User{
		{Name: "Admin", Email: "admin@example.com", Password: "hashed_password"},
		{Name: "John Doe", Email: "john@example.com", Password: "hashed_password"},
	}

	for _, u := range users {
		if err := db.FirstOrCreate(&u, User{Email: u.Email}).Error; err != nil {
			return err
		}
	}

	return nil
}
package db

import (
	"log"

	"gorm.io/gorm"

	"ajari/service-users/internal/db/migration"
)

// Migrate runs all database migrations in order.
func Migrate(db *gorm.DB) {
	log.Println("🚀 Starting database migrations...")

	if err := migration.CreateUsersTable(db); err != nil {
		log.Fatalf("Migration CreateUsersTable failed: %v", err)
	}

	log.Println("✅ Database migrations completed successfully!")
}
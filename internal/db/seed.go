package db

import (
	"log"

	"gorm.io/gorm"

	"ajari/service-users/internal/db/seed"
)

// Seed runs all database seeders.
func Seed(db *gorm.DB) {
	log.Println("🌱 Starting database seeding...")

	if err := seed.Users(db); err != nil {
		log.Fatalf("SeedUsers failed: %v", err)
	}

	log.Println("✅ Database seeding completed successfully!")
}
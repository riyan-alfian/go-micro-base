package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserActivity represents a log of what a user did in the system.
type UserActivity struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	ActivityType string             `bson:"activity_type" json:"activity_type"` // e.g. "login", "watch_video", "complete_exam"
	TargetID     *primitive.ObjectID `bson:"target_id,omitempty" json:"target_id,omitempty"` // e.g. video ID, course ID, etc.
	Description  string              `bson:"description,omitempty" json:"description,omitempty"`
	Metadata     map[string]any      `bson:"metadata,omitempty" json:"metadata,omitempty"`   // additional info like duration, device, etc.
	CreatedAt    time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time           `bson:"updated_at" json:"updated_at"`
}

// CollectionName returns the MongoDB collection name.
func (UserActivity) CollectionName() string {
	return "user_activities"
}

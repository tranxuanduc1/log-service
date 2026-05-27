package logactivity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type LogActivity struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey"            json:"id"`
	ActionType string         `gorm:"type:varchar(100);not null"      json:"action_type"`
	IPAddress  *string        `gorm:"type:inet"                       json:"ip_address,omitempty"`
	DurationMS *int           `gorm:"type:integer"                    json:"duration_ms,omitempty"`
	Metadata   datatypes.JSON `gorm:"type:jsonb"                      json:"metadata,omitempty"`
	CreatedAt  time.Time      `gorm:"type:timestamptz;autoCreateTime" json:"created_at"`
	UserID     *uuid.UUID     `gorm:"type:uuid"                       json:"user_id,omitempty"`
}

// BeforeCreate auto-generates a UUIDv7 for every new record
func (l *LogActivity) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	l.ID = id
	return nil
}

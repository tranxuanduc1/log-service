package logactivity

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type CreateLogActivityRequest struct {
	ActionType string         `json:"action_type" binding:"required"`
	IPAddress  *string        `json:"ip_address,omitempty"`
	DurationMS *int           `json:"duration_ms,omitempty"`
	Metadata   datatypes.JSON `json:"metadata,omitempty"`
	UserID     *uuid.UUID     `json:"user_id,omitempty"`
}

type GetLogActivityResponse struct {
	ID         uuid.UUID      `json:"id"`
	ActionType string         `json:"action_type"`
	IPAddress  *string        `json:"ip_address,omitempty"`
	DurationMS *int           `json:"duration_ms,omitempty"`
	Metadata   datatypes.JSON `json:"metadata,omitempty"`
	CreatedAt  string         `json:"created_at"` // ISO8601 format
	UserID     *uuid.UUID     `json:"user_id,omitempty"`
}

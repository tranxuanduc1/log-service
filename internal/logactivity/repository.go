package logactivity

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, log *LogActivity) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *Repository) FindAll(ctx context.Context) ([]LogActivity, error) {
	var logs []LogActivity

	err := r.db.
		WithContext(ctx).
		Order("created_at DESC").
		Find(&logs).
		Error

	return logs, err
}

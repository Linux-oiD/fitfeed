package profiledb

import (
	"context"
	"fitfeed/api/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProfileDB {
	return &ProfileDB{db: db}
}

func (p *ProfileDB) GetByUserID(ctx context.Context, userID uuid.UUID) (entity.Profile, error) {
	profile, err := gorm.G[entity.Profile](p.db).Where("user_id = ?", userID).First(ctx)
	return profile, err
}

func (p *ProfileDB) Update(ctx context.Context, userID uuid.UUID, profile entity.Profile) error {
	_, err := gorm.G[entity.Profile](p.db).Where("user_id = ?", userID).Updates(ctx, profile)
	return err
}

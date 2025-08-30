package profiledb

import (
	"context"

	"fitfeed/auth/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProfileDB {

	return &ProfileDB{db: db}
}

func (p *ProfileDB) GetByID(ctx context.Context, id uuid.UUID) (entity.Profile, error) {

	profile, err := gorm.G[entity.Profile](p.db).Where("id = ?", id).First(ctx)
	return profile, err
}

func (p *ProfileDB) GetByEmail(ctx context.Context, email string) (entity.Profile, error) {

	profile, err := gorm.G[entity.Profile](p.db).Where("email = ?", email).First(ctx)
	return profile, err
}

func (p *ProfileDB) Update(ctx context.Context, id uuid.UUID, profileUpd entity.Profile) error {

	_, err := gorm.G[entity.Profile](p.db).Where("id = ?", id).Updates(ctx, profileUpd)
	return err
}

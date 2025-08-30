package oauthdb

import (
	"context"

	"fitfeed/auth/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OauthDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *OauthDB {

	return &OauthDB{db: db}
}

func (o *OauthDB) Create(ctx context.Context, provider entity.OauthProvider) error {

	err := gorm.G[entity.OauthProvider](o.db).Create(ctx, &provider)
	return err
}

func (o *OauthDB) GetByID(ctx context.Context, id uuid.UUID) (entity.OauthProvider, error) {

	provider, err := gorm.G[entity.OauthProvider](o.db).Where("id = ?", id).First(ctx)
	return provider, err
}

func (o *OauthDB) GetByProviderID(ctx context.Context, providerID string) (entity.OauthProvider, error) {

	provider, err := gorm.G[entity.OauthProvider](o.db).Where("provider_id = ?", providerID).First(ctx)
	return provider, err
}

func (o *OauthDB) UpdateProviderID(ctx context.Context, id uuid.UUID, providerID string) error {

	_, err := gorm.G[entity.OauthProvider](o.db).Where("id = ?", id).Update(ctx, "provider_id", providerID)
	return err
}

func (o *OauthDB) Delete(ctx context.Context, id uuid.UUID) error {

	_, err := gorm.G[entity.OauthProvider](o.db).Where("id = ?", id).Delete(ctx)
	return err
}

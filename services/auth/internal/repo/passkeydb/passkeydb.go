package passkeydb

import (
	"context"
	"fitfeed/auth/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PasskeyDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PasskeyDB {
	return &PasskeyDB{db: db}
}

func (p *PasskeyDB) Create(ctx context.Context, pk entity.Passkey) error {
	err := gorm.G[entity.Passkey](p.db).Create(ctx, &pk)
	return err
}

func (p *PasskeyDB) GetByCredentialID(ctx context.Context, credentialID []byte) (entity.Passkey, error) {
	pk, err := gorm.G[entity.Passkey](p.db).Where("credential_id = ?", credentialID).First(ctx)
	return pk, err
}

func (p *PasskeyDB) UpdateSignCount(ctx context.Context, credentialID []byte, signCount uint32) error {
	_, err := gorm.G[entity.Passkey](p.db).Where("credential_id = ?", credentialID).Update(ctx, "sign_count", signCount)
	return err
}

func (p *PasskeyDB) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := gorm.G[entity.Passkey](p.db).Where("id = ?", id).Delete(ctx)
	return err
}

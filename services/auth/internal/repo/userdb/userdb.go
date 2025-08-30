package userdb

import (
	"context"

	"fitfeed/auth/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) Create(ctx context.Context, user entity.User) error {

	err := gorm.G[entity.User](u.db).Create(ctx, &user)
	return err
}

func (u *UserDB) GetByID(ctx context.Context, id uuid.UUID) (entity.User, error) {

	user, err := gorm.G[entity.User](u.db).Where("id = ?", id).First(ctx)
	return user, err
}

func (u *UserDB) GetByUsername(ctx context.Context, username string) (entity.User, error) {

	user, err := gorm.G[entity.User](u.db).Where("username = ?", username).First(ctx)
	return user, err
}

func (u *UserDB) UpdateUsername(ctx context.Context, id uuid.UUID, username string) error {

	_, err := gorm.G[entity.User](u.db).Where("id = ?", id).Update(ctx, "username", username)
	return err

}

func (u *UserDB) Delete(ctx context.Context, id uuid.UUID) error {

	_, err := gorm.G[entity.User](u.db).Where("id = ?", id).Delete(ctx)
	return err
}

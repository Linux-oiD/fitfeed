package usermanager

import (
	"context"
	"fitfeed/auth/internal/entity"
	"log/slog"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Mocks
type mockUserDB struct {
	mock.Mock
}

func (m *mockUserDB) Create(ctx context.Context, u entity.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *mockUserDB) GetByID(ctx context.Context, id uuid.UUID) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *mockUserDB) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *mockUserDB) UpdateUsername(ctx context.Context, id uuid.UUID, username string) error {
	args := m.Called(ctx, id, username)
	return args.Error(0)
}

func (m *mockUserDB) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type mockProfileDB struct {
	mock.Mock
}

func (m *mockProfileDB) Create(ctx context.Context, p entity.Profile) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *mockProfileDB) GetByID(ctx context.Context, id uuid.UUID) (entity.Profile, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.Profile), args.Error(1)
}

func (m *mockProfileDB) GetByEmail(ctx context.Context, email string) (entity.Profile, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(entity.Profile), args.Error(1)
}

func (m *mockProfileDB) Update(ctx context.Context, id uuid.UUID, p entity.Profile) error {
	args := m.Called(ctx, id, p)
	return args.Error(0)
}

func TestCheckUsername(t *testing.T) {
	udb := new(mockUserDB)
	pdb := new(mockProfileDB)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	um := New(udb, pdb, logger)

	t.Run("available", func(t *testing.T) {
		udb.On("GetByUsername", mock.Anything, "newuser").Return(entity.User{}, gorm.ErrRecordNotFound).Once()
		err := um.CheckUsername(context.Background(), "newuser")
		assert.NoError(t, err)
	})

	t.Run("taken", func(t *testing.T) {
		udb.On("GetByUsername", mock.Anything, "taken").Return(entity.User{Username: "taken"}, nil).Once()
		err := um.CheckUsername(context.Background(), "taken")
		assert.Equal(t, entity.ENOTAVAILABLE, err)
	})
}

func TestRegisterUser(t *testing.T) {
	udb := new(mockUserDB)
	pdb := new(mockProfileDB)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	um := New(udb, pdb, logger)

	t.Run("success", func(t *testing.T) {
		user := entity.User{Username: "test"}
		udb.On("Create", mock.Anything, user).Return(nil).Once()
		pdb.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		err := um.RegisterUser(context.Background(), user)
		assert.NoError(t, err)
	})
}

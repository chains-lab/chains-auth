package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/chains-lab/chains-auth/internal/config"
	"github.com/chains-lab/chains-auth/internal/repo/sqldb"
	"github.com/chains-lab/gatekit/roles"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserModel struct {
	ID           uuid.UUID  `json:"id"`
	Email        string     `json:"email"`
	Role         roles.Role `json:"role"`
	Subscription uuid.UUID  `json:"subscription"`
	Verified     bool       `json:"verified,omitempty"`
	Suspended    bool       `json:"suspended,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at"`
	CreatedAt    time.Time  `json:"created_at"`
}

type usersSQL interface {
	New() sqldb.UserQ
	Insert(ctx context.Context, input sqldb.UserModel) error
	Delete(ctx context.Context) error
	Count(ctx context.Context) (int, error)
	Select(ctx context.Context) ([]sqldb.UserModel, error)
	Get(ctx context.Context) (sqldb.UserModel, error)

	FilterID(id uuid.UUID) sqldb.UserQ
	FilterEmail(email string) sqldb.UserQ
	FilterRole(role roles.Role) sqldb.UserQ
	FilterSubscription(subscription uuid.UUID) sqldb.UserQ
	FilterVerified(verified bool) sqldb.UserQ

	Update(ctx context.Context, input sqldb.UserUpdateInput) error

	Page(limit, offset uint64) sqldb.UserQ
	Transaction(fn func(ctx context.Context) error) error

	Drop(ctx context.Context) error
}

type UsersRepo struct {
	sql usersSQL
	log *logrus.Entry
}

func NewUsers(cfg config.Config, log *logrus.Logger) (UsersRepo, error) {
	db, err := sql.Open("postgres", cfg.Database.SQL.URL)
	if err != nil {
		return UsersRepo{}, err
	}

	sqlImpl := sqldb.NewUsers(db)

	return UsersRepo{
		sql: sqlImpl,
		log: log.WithField("component", "users"),
	}, nil
}

func (a UsersRepo) Create(ctx context.Context, input UserModel) error {
	ctxSync, cancel := context.WithTimeout(ctx, dataCtxTimeAisle)
	defer cancel()

	if err := a.sql.Insert(ctxSync, sqldb.UserModel{
		ID:           input.ID,
		Email:        input.Email,
		Role:         input.Role,
		Subscription: input.Subscription,
		Verified:     input.Verified,
		Suspended:    input.Suspended,
		UpdatedAt:    input.UpdatedAt,
		CreatedAt:    input.CreatedAt,
	}); err != nil {
		return err
	}

	return nil
}

type UserUpdateRequest struct {
	Role         *roles.Role `json:"role"`
	Subscription *uuid.UUID  `json:"subscription"`
	Verified     *bool       `json:"verified,omitempty"`
	Suspended    *bool       `json:"suspended,omitempty"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

func (a UsersRepo) Update(ctx context.Context, ID uuid.UUID, input UserUpdateRequest) error {
	ctxSync, cancel := context.WithTimeout(ctx, dataCtxTimeAisle)
	defer cancel()

	var sqlInput sqldb.UserUpdateInput
	if input.Role != nil {
		sqlInput.Role = input.Role
	}
	if input.Subscription != nil {
		sqlInput.Subscription = input.Subscription
	}
	if input.Verified != nil {
		sqlInput.Verified = input.Verified
	}
	if input.Suspended != nil {
		sqlInput.Suspended = input.Suspended
	}
	sqlInput.UpdatedAt = input.UpdatedAt

	if err := a.sql.New().FilterID(ID).Update(ctxSync, sqlInput); err != nil {
		return err
	}

	return nil
}

func (a UsersRepo) Delete(ctx context.Context, ID uuid.UUID) error {
	ctxSync, cancel := context.WithTimeout(ctx, dataCtxTimeAisle)
	defer cancel()

	if err := a.sql.New().FilterID(ID).Delete(ctxSync); err != nil {
		return err
	}

	return nil
}

func (a UsersRepo) GetByID(ctx context.Context, ID uuid.UUID) (UserModel, error) {
	ctxSync, cancel := context.WithTimeout(ctx, dataCtxTimeAisle)
	defer cancel()

	user, err := a.sql.New().FilterID(ID).Get(ctxSync)
	if err != nil {
		return UserModel{}, err
	}

	res := UserModel{
		ID:           user.ID,
		Email:        user.Email,
		Role:         user.Role,
		Subscription: user.Subscription,
		Verified:     user.Verified,
		Suspended:    user.Suspended,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return res, nil
}

func (a UsersRepo) GetByEmail(ctx context.Context, email string) (UserModel, error) {
	ctxSync, cancel := context.WithTimeout(context.Background(), dataCtxTimeAisle)
	defer cancel()

	user, err := a.sql.New().FilterEmail(email).Get(ctxSync)
	if err != nil {
		return UserModel{}, err
	}

	res := UserModel{
		ID:           user.ID,
		Email:        user.Email,
		Role:         user.Role,
		Subscription: user.Subscription,
		Verified:     user.Verified,
		Suspended:    user.Suspended,
		UpdatedAt:    user.UpdatedAt,
		CreatedAt:    user.CreatedAt,
	}

	return res, nil
}

func (a UsersRepo) Transaction(fn func(ctx context.Context) error) error {
	return a.sql.Transaction(fn)
}

func (a UsersRepo) Drop(ctx context.Context) error {
	ctxSync, cancel := context.WithTimeout(ctx, dataCtxTimeAisle)
	defer cancel()

	if err := a.sql.Drop(ctxSync); err != nil {
		return err
	}

	return nil
}

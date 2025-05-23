package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/chains-lab/chains-auth/internal/config"
	"github.com/chains-lab/chains-auth/internal/repo/redisdb"
	"github.com/chains-lab/chains-auth/internal/repo/sqldb"
	"github.com/chains-lab/gatekit/roles"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Account struct {
	ID           uuid.UUID  `json:"id"`
	Email        string     `json:"email"`
	Role         roles.Role `json:"role"`
	Subscription uuid.UUID  `json:"subscription,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at"`
	CreatedAt    time.Time  `json:"created_at"`
}

type accountSQL interface {
	New() sqldb.AccountQ
	Insert(ctx context.Context, input sqldb.AccountInsertInput) error
	Delete(ctx context.Context) error
	Count(ctx context.Context) (int, error)
	Select(ctx context.Context) ([]sqldb.AccountModel, error)
	Get(ctx context.Context) (sqldb.AccountModel, error)

	FilterID(id uuid.UUID) sqldb.AccountQ
	FilterEmail(email string) sqldb.AccountQ
	FilterRole(role roles.Role) sqldb.AccountQ
	FilterSubscription(subscription uuid.UUID) sqldb.AccountQ

	Update(ctx context.Context, input sqldb.AccountUpdateInput) error

	Page(limit, offset uint64) sqldb.AccountQ
	Transaction(fn func(ctx context.Context) error) error

	Drop(ctx context.Context) error
}

type accountRedis interface {
	Create(ctx context.Context, input redisdb.CreateAccountInput) error
	Set(ctx context.Context, input redisdb.AccountSetInput) error
	Update(ctx context.Context, accountID uuid.UUID, input redisdb.AccountUpdateRequest) error
	GetByID(ctx context.Context, accountID string) (redisdb.AccountModel, error)
	GetByEmail(ctx context.Context, email string) (redisdb.AccountModel, error)
	Delete(ctx context.Context, accountID, email string) error

	Drop(ctx context.Context) error
}

type AccountsRepo struct {
	sql   sqldb.AccountQ
	redis accountRedis
	log   *logrus.Entry
}

func NewAccounts(cfg config.Config, log *logrus.Logger) (AccountsRepo, error) {
	db, err := sql.Open("postgres", cfg.Database.SQL.URL)
	if err != nil {
		return AccountsRepo{}, err
	}

	redisImpl := redisdb.NewAccounts(redis.NewClient(&redis.Options{
		Addr:     cfg.Database.Redis.Addr,
		Password: cfg.Database.Redis.Password,
		DB:       cfg.Database.Redis.DB,
	}), cfg.Database.Redis.Lifetime)
	sqlImpl := sqldb.NewAccounts(db)

	return AccountsRepo{
		sql:   sqlImpl,
		redis: redisImpl,
		log:   log.WithField("component", "accounts"),
	}, nil
}

type AccountCreateRequest struct {
	ID           uuid.UUID  `json:"id"`
	Email        string     `json:"email"`
	Role         roles.Role `json:"role"`
	Subscription uuid.UUID  `json:"subscription,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}

func (a AccountsRepo) Create(ctx context.Context, input AccountCreateRequest) error {
	ctxSync, cancel := context.WithTimeout(ctx, dataCtxTimeAisle)
	defer cancel()

	if err := a.sql.Insert(ctxSync, sqldb.AccountInsertInput{
		ID:           input.ID,
		Email:        input.Email,
		Role:         input.Role,
		Subscription: input.Subscription,
		CreatedAt:    input.CreatedAt,
	}); err != nil {
		return err
	}

	_ = a.redis.Create(ctxSync, redisdb.CreateAccountInput{
		ID:           input.ID,
		Email:        input.Email,
		Role:         input.Role,
		Subscription: input.Subscription,
	})

	return nil
}

type AccountUpdateRequest struct {
	Role         *roles.Role `json:"role"`
	Subscription *uuid.UUID  `json:"subscription,omitempty"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

func (a AccountsRepo) Update(ctx context.Context, ID uuid.UUID, input AccountUpdateRequest) error {
	ctxSync, cancel := context.WithTimeout(ctx, dataCtxTimeAisle)
	defer cancel()

	var sqlInput sqldb.AccountUpdateInput
	if input.Role != nil {
		sqlInput.Role = input.Role
	}
	if input.Subscription != nil {
		sqlInput.Subscription = input.Subscription
	}
	sqlInput.UpdatedAt = input.UpdatedAt

	if err := a.sql.New().FilterID(ID).Update(ctxSync, sqldb.AccountUpdateInput{
		Role:         input.Role,
		Subscription: input.Subscription,
		UpdatedAt:    input.UpdatedAt,
	}); err != nil {
		return err
	}

	account, err := a.sql.New().FilterID(ID).Get(ctxSync)
	if err != nil {
		return err
	}

	_ = a.redis.Set(ctxSync, redisdb.AccountSetInput{
		ID:           account.ID,
		Email:        account.Email,
		Role:         account.Role,
		Subscription: account.Subscription,
		UpdatedAt:    account.UpdatedAt,
		CreatedAt:    account.CreatedAt,
	})
	return nil
}

func (a AccountsRepo) Delete(ctx context.Context, ID uuid.UUID) error {
	ctxSync, cancel := context.WithTimeout(ctx, dataCtxTimeAisle)
	defer cancel()

	account, err := a.sql.New().FilterID(ID).Get(ctxSync)
	if err != nil {
		return err
	}

	if err := a.redis.Delete(ctxSync, account.ID.String(), account.Email); err != nil {
		a.log.WithField("database", "redis").Errorf("error creating account in redis: %v", err)
	}

	if err := a.sql.New().FilterID(ID).Delete(ctxSync); err != nil {
		return err
	}

	return nil
}

func (a AccountsRepo) GetByID(ctx context.Context, ID uuid.UUID) (Account, error) {
	ctxSync, cancel := context.WithTimeout(ctx, dataCtxTimeAisle)
	defer cancel()

	redisRes, err := a.redis.GetByID(ctxSync, ID.String())
	if err != nil {
		a.log.WithField("database", "redis").Errorf("error creating account in redis: %v", err)
	} else {
		res := Account{
			ID:           redisRes.ID,
			Email:        redisRes.Email,
			Subscription: redisRes.Subscription,
			CreatedAt:    redisRes.CreatedAt,
			Role:         redisRes.Role,
		}
		if redisRes.UpdatedAt != nil {
			res.UpdatedAt = *redisRes.UpdatedAt
		}
		return res, nil
	}

	account, err := a.sql.New().FilterID(ID).Get(ctxSync)
	if err != nil {
		return Account{}, err
	}

	res := Account{
		ID:           account.ID,
		Email:        account.Email,
		Role:         account.Role,
		Subscription: account.Subscription,
		CreatedAt:    account.CreatedAt,
		UpdatedAt:    *account.UpdatedAt,
	}
	if account.UpdatedAt != nil {
		res.UpdatedAt = *account.UpdatedAt
	}

	if err := a.redis.Set(ctxSync, redisdb.AccountSetInput{
		ID:           account.ID,
		Email:        account.Email,
		Role:         account.Role,
		Subscription: account.Subscription,
		UpdatedAt:    account.UpdatedAt,
		CreatedAt:    account.CreatedAt,
	}); err != nil {
		a.log.WithField("database", "redis").Errorf("error creating account in redis: %v", err)
	}

	return res, nil
}

func (a AccountsRepo) GetByEmail(ctx context.Context, email string) (Account, error) {
	ctxSync, cancel := context.WithTimeout(context.Background(), dataCtxTimeAisle)
	defer cancel()

	redisRes, err := a.redis.GetByEmail(ctxSync, email)
	if err != nil {
		a.log.WithField("database", "redis").Errorf("error creating account in redis: %v", err)
	} else {
		res := Account{
			ID:           redisRes.ID,
			Email:        redisRes.Email,
			Subscription: redisRes.Subscription,
			CreatedAt:    redisRes.CreatedAt,
			Role:         redisRes.Role,
		}
		if redisRes.UpdatedAt != nil {
			res.UpdatedAt = *redisRes.UpdatedAt
		}
		return res, nil
	}

	account, err := a.sql.New().FilterEmail(email).Get(ctxSync)
	if err != nil {
		return Account{}, err
	}

	res := Account{
		ID:           account.ID,
		Email:        account.Email,
		Role:         account.Role,
		Subscription: account.Subscription,
		CreatedAt:    account.CreatedAt,
	}

	if account.UpdatedAt != nil {
		res.UpdatedAt = *account.UpdatedAt
	}

	if err := a.redis.Set(ctxSync, redisdb.AccountSetInput{
		ID:           account.ID,
		Email:        account.Email,
		Role:         account.Role,
		Subscription: account.Subscription,
		UpdatedAt:    account.UpdatedAt,
		CreatedAt:    account.CreatedAt,
	}); err != nil {
		a.log.WithField("database", "redis").Errorf("error creating account in redis: %v", err)
	}

	return res, nil
}

func (a AccountsRepo) Transaction(fn func(ctx context.Context) error) error {
	return a.sql.Transaction(fn)
}

func (a AccountsRepo) Drop(ctx context.Context) error {
	ctxSync, cancel := context.WithTimeout(ctx, dataCtxTimeAisle)
	defer cancel()

	if err := a.sql.Drop(ctxSync); err != nil {
		return err
	}

	if err := a.redis.Drop(ctxSync); err != nil {
		return err
	}

	return nil
}

package repository

import (
	"back-end-hr-manajement/domain"
	"back-end-hr-manajement/dto"
	"context"

	"github.com/doug-martin/goqu/v9"
)

type userRepository struct {
	db *goqu.Database
}

func NewUserRepository(db *goqu.Database) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

// DeleteTokenRefresh implements [domain.UserRepository].
func (u *userRepository) DeleteTokenRefresh(ctx context.Context, userId string) error {
	dataset := u.db.Delete("users").Where(goqu.C("id").Eq(userId)).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// FindToken implements [domain.UserRepository].
func (u *userRepository) FindToken(ctx context.Context, userId string) (result dto.TokenResponse, err error) {
	dataset := u.db.From("refresh_tokens").Where(goqu.C("user_id").Eq(userId))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// FindtUserByEmail implements [domain.UserRepository].
func (u *userRepository) FindtUserByEmail(ctx context.Context, email string) (result domain.UserDomain, err error) {
	dataset := u.db.From("users").Where(goqu.C("email").Eq(email))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// SaveUser implements [domain.UserRepository].
func (u *userRepository) SaveUser(ctx context.Context, user domain.UserDomain) error {
	dataset := u.db.Insert("users").Rows(user).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// SaveTokenRefresh implements [domain.UserRepository].
func (u *userRepository) SaveTokenRefresh(ctx context.Context, tokenData domain.RefreshTokenDomain) error {
	dataset := u.db.Insert("refresh_tokens").Rows(tokenData).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// UpdateTokenRefresh implements [domain.UserRepository].
func (u *userRepository) UpdateTokenRefresh(ctx context.Context, userId string, tokenData string, time string) error {
	ds := u.db.
		Update(goqu.T("refresh_tokens")).
		Set(goqu.Record{
			"token":      tokenData,
			"updated_at": time,
		}).
		Where(goqu.C("user_id").Eq(userId))

	_, err := ds.Executor().ExecContext(ctx)
	return err
}

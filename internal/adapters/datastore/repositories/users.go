/* Copyright (C) Fedir Petryk */

package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
)

type UserRepository struct {
	db *bun.DB
}

func NewUser(db *bun.DB) UserRepository {
	return UserRepository{db: db}
}

// GetByEmail retrieves a user with the given email address from the database.
func (r UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return user, nil
}

// Create creates a new user in the database.
func (r UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// Delete deletes a user with the given ID from the database.
func (r UserRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().Model(&domain.User{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("user with ID %d not deleted: %s", id, err)
	}

	return nil
}

// Update updates an existing user in the database.
func (r UserRepository) Update(ctx context.Context, user *domain.User) error {
	_, err := r.db.NewUpdate().Model(user).OmitZero().WherePK().Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// UpdateActivity deactivates an existing user in the database.
func (r UserRepository) UpdateActivity(ctx context.Context, userID int, active bool) error {
	_, err := r.db.NewUpdate().
		Model(&domain.User{}).
		Set("active = ?", active).
		Where("id = ?", userID).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// SetNewPwd sets new password for user
func (r UserRepository) SetNewPwd(ctx context.Context, user *domain.User) error {
	_, err := r.db.NewUpdate().Model(user).Column("password").WherePK().Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Get retrieves a user with the given ID from the database.
func (r UserRepository) Get(ctx context.Context, id int) (*domain.User, error) {
	user := &domain.User{ID: id}
	err := r.db.NewSelect().Model(user).WherePK().Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

// List retrieves a list of users from the database.
func (r UserRepository) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.NewSelect().Model(&users).Limit(limit).Offset(offset).Order("created_at DESC").Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	return users, nil
}

// Create creates a new user in the database.
func (r UserRepository) Create(ctx context.Context, user domain.User, token domain.UserTokenAction) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	_, err = tx.NewInsert().Model(&user).Returning("id").Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	token.UserID = &user.ID

	_, err = tx.NewInsert().Model(&token).Exec(ctx)
	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return errors.Wrap(
				errors.Wrap(errTx, "failed to rollback transaction"), "failed to create verification token",
			)
		}

		return errors.Wrap(err, "failed to create verification token")
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// EmailExists check if email exists in db
func (r UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	count, err := r.db.NewSelect().Model(&domain.User{}).Where("email = ?", email).ScanAndCount(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("failed to get user by email: %w", err)
	}

	return count > 0, nil
}

package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/adhitamafikri/cozy-prop-tech/backend/api/internal/domain/user"
	cerrors "github.com/adhitamafikri/cozy-prop-tech/backend/api/internal/errors"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	query := `
		INSERT INTO users (name, email, phone, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`

	row := r.db.QueryRowContext(ctx, query, u.Name, u.Email, u.Phone, u.Password)

	var newUser user.User
	err := row.Scan(&newUser.ID, &newUser.CreatedAt, &newUser.UpdatedAt)
	if err != nil {
		return err
	}

	u.ID = newUser.ID
	u.CreatedAt = newUser.CreatedAt
	u.UpdatedAt = newUser.UpdatedAt

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*user.User, error) {
	query := `
		SELECT id, name, email, phone, password, created_at, updated_at, deleted_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`

	var u user.User
	err := r.db.GetContext(ctx, &u, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, cerrors.ErrUserNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `
		SELECT id, name, email, phone, password, created_at, updated_at, deleted_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`

	var u user.User
	err := r.db.GetContext(ctx, &u, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, cerrors.ErrUserNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) GetByPhone(ctx context.Context, phone string) (*user.User, error) {
	query := `
		SELECT id, name, email, phone, password, created_at, updated_at, deleted_at
		FROM users
		WHERE phone = $1 AND deleted_at IS NULL
	`

	var u user.User
	err := r.db.GetContext(ctx, &u, query, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, cerrors.ErrUserNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) Update(ctx context.Context, u *user.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, phone = $3, updated_at = NOW()
		WHERE id = $4 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, u.Name, u.Email, u.Phone, u.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return cerrors.ErrUserNotFound
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `
		UPDATE users
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return cerrors.ErrUserNotFound
	}

	return nil
}

func (r *UserRepository) List(ctx context.Context, limit, offset int) ([]user.User, error) {
	query := `
		SELECT id, name, email, phone, password, created_at, updated_at, deleted_at
		FROM users
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	var users []user.User
	err := r.db.SelectContext(ctx, &users, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// RBAC - Role management

func (r *UserRepository) AssignRole(ctx context.Context, userID, roleID int64) error {
	query := `
		INSERT INTO user_roles (user_id, role_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id, role_id) DO NOTHING
	`

	_, err := r.db.ExecContext(ctx, query, userID, roleID)
	return err
}

func (r *UserRepository) RemoveRole(ctx context.Context, userID, roleID int64) error {
	query := `
		DELETE FROM user_roles
		WHERE user_id = $1 AND role_id = $2
	`

	result, err := r.db.ExecContext(ctx, query, userID, roleID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return cerrors.ErrNotFound
	}

	return nil
}

func (r *UserRepository) GetRoles(ctx context.Context, userID int64) ([]user.Role, error) {
	query := `
		SELECT r.id, r.name, r.description, r.created_at, r.updated_at, r.deleted_at
		FROM roles r
		INNER JOIN user_roles ur ON r.id = ur.role_id
		WHERE ur.user_id = $1 AND r.deleted_at IS NULL
	`

	var roles []user.Role
	err := r.db.SelectContext(ctx, &roles, query, userID)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// RBAC - Permission management

func (r *UserRepository) HasPermission(ctx context.Context, userID int64, permissionName string) (bool, error) {
	query := `
		SELECT COUNT(1)
		FROM users u
		INNER JOIN user_roles ur ON u.id = ur.user_id
		INNER JOIN role_permissions rp ON ur.role_id = rp.role_id
		INNER JOIN permissions p ON rp.permission_id = p.id
		WHERE u.id = $1 
			AND u.deleted_at IS NULL
			AND p.name = $2 
			AND p.deleted_at IS NULL
	`

	var count int
	err := r.db.GetContext(ctx, &count, query, userID, permissionName)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

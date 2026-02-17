package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByPhone(ctx context.Context, phone string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]User, error)

	// RBAC - Role management
	AssignRole(ctx context.Context, userID, roleID int64) error
	RemoveRole(ctx context.Context, userID, roleID int64) error
	GetRoles(ctx context.Context, userID int64) ([]Role, error)

	// RBAC - Permission management
	HasPermission(ctx context.Context, userID int64, permissionName string) (bool, error)
}

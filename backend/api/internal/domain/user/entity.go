package user

import "time"

type User struct {
	ID        int64      `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	Phone     string     `json:"phone" db:"phone"`
	Password  string     `json:"-" db:"password"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Role struct {
	ID          int64      `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Permission struct {
	ID          int64      `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type UserRole struct {
	UserID int64 `json:"user_id" db:"user_id"`
	RoleID int64 `json:"role_id" db:"role_id"`
}

type RolePermission struct {
	RoleID       int64 `json:"role_id" db:"role_id"`
	PermissionID int64 `json:"permission_id" db:"permission_id"`
}

type UserWithRoles struct {
	User
	Roles []Role `json:"roles,omitempty"`
}

type RoleWithPermissions struct {
	Role
	Permissions []Permission `json:"permissions,omitempty"`
}

package auth

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository interface
type Repository interface {
	GetUsers() ([]*UserWithoutPassword, error)
	Login(login Login) (*UserWithPassword, error)
	Register(register Register) (*string, error)

	// User Role
	AddUserRole(userRole UserRoleInput) error
	DeleteUserRole(userRole UserRoleInput) error

	// Roles
	GetRoles() ([]*RoleOutput, error)
	GetRoleByID(id string) (*RoleOutput, error)
	AddRole(role RoleInput) error
	DeleteRoleByID(is string) error
}

// repo struct
type repo struct {
	client *mongo.Client
}

// NewRepo user repository
func NewRepo(url string) Repository {
	clientOptions := options.Client().ApplyURI(url)
	client, _ := mongo.NewClient(clientOptions)
	return &repo{client: client}
}

// GetUsers method
func (r *repo) GetUsers() ([]*UserWithoutPassword, error) {
	return nil, nil
}

// Login method
func (r *repo) Login(login Login) (*UserWithPassword, error) {
	return nil, nil
}

// Register method
func (r *repo) Register(register Register) (*string, error) {
	return nil, nil
}

// AddUserRole method
func (r *repo) AddUserRole(userRole UserRoleInput) error {
	return nil
}

// DeleteUserRole method
func (r *repo) DeleteUserRole(userRole UserRoleInput) error {
	return nil
}

// GetRoles method
func (r *repo) GetRoles() ([]*RoleOutput, error) {
	return nil, nil
}

// GetRoleByID method
func (r *repo) GetRoleByID(id string) (*RoleOutput, error) {
	return nil, nil
}

// AddRole method
func (r *repo) AddRole(role RoleInput) error {
	return nil
}

// DeleteRoleByID method
func (r *repo) DeleteRoleByID(is string) error {
	return nil
}

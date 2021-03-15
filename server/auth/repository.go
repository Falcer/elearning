package auth

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Repository interface
type Repository interface {
	GetUsers() (*[]UserWithoutPassword, error)
	Login(email string) (*UserWithPassword, error)
	Register(register Register) (*string, error)

	// User Role
	AddUserRole(userRole UserRoleInput) error
	DeleteUserRole(userRole UserRoleInput) error

	// Roles
	GetRoles() (*[]RoleOutput, error)
	GetRoleByID(id string) (*RoleOutput, error)
	AddRole(role RoleInput) error
	DeleteRoleByID(is string) error
}

var (
	databaseName string
)

// repo struct
type repo struct {
	client *mongo.Client
}

// NewRepo user repository
func NewRepo(url string) Repository {
	databaseName = os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		databaseName = "elearning"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil
	}
	return &repo{client: client}
}

// GetUsers method
func (r *repo) GetUsers() (*[]UserWithoutPassword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := r.client.Database(databaseName).Collection("users").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var users []UserWithoutPassword
	for cur.Next(ctx) {
		var result UserWithoutPassword
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		users = append(users, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

// Login method
func (r *repo) Login(email string) (*UserWithPassword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur := r.client.Database(databaseName).Collection("users").FindOne(ctx, bson.D{})
	var user UserWithPassword
	cur.Decode(&user)
	return &user, nil
}

// Register method
func (r *repo) Register(register Register) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id, err := r.client.Database(databaseName).Collection("users").InsertOne(ctx, register)
	if err != nil {
		return nil, err
	}
	resID := id.InsertedID.(primitive.ObjectID).Hex()
	return &resID, nil
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
func (r *repo) GetRoles() (*[]RoleOutput, error) {
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

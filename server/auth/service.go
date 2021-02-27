package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Service interface
type Service interface {
	// Auth
	GetUsers() ([]*UserWithoutPassword, error)
	Login(login Login) (*UserWithoutPassword, error)
	Register(register Register) (*UserWithoutPassword, error)
	RefreshToken(refeshToken string) (*UserToken, error)
	Verify(tokenString string) error

	// User Role
	AddUserRole(userRole UserRoleInput) (*UserWithRole, error)
	DeleteUserRole(userRole UserRoleInput) error

	// Roles
	GetRoles() ([]*RoleOutput, error)
	GetRoleByID(id string) (*RoleOutput, error)
	AddRole(role RoleInput) (*RoleOutput, error)
	DeleteRoleByID(is string) error
}

type service struct {
	repo Repository
}

// NewService service
func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetUsers() ([]*UserWithoutPassword, error) {
	return nil, nil
}

func (s *service) Login(login Login) (*UserWithoutPassword, error) {
	return nil, nil
}

func (s *service) Register(register Register) (*UserWithoutPassword, error) {
	return nil, nil
}

func (s *service) AddUserRole(userRole UserRoleInput) (*UserWithRole, error) {
	return nil, nil
}

func (s *service) DeleteUserRole(userRole UserRoleInput) error {
	return nil
}

func (s *service) GetRoles() ([]*RoleOutput, error) {
	return nil, nil
}

func (s *service) GetRoleByID(id string) (*RoleOutput, error) {
	return nil, nil
}

func (s *service) AddRole(role RoleInput) (*RoleOutput, error) {
	return nil, nil
}

func (s *service) DeleteRoleByID(is string) error {
	return nil
}

func (s *service) RefreshToken(refeshToken string) (*UserToken, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refeshToken, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error : %s", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("Token not valid")
	}
	newToken, err := createToken(claims["id"].(string))
	return newToken, err
}

func (s *service) Verify(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return fmt.Errorf("Error : %s", err)
	}
	if !token.Valid {
		return fmt.Errorf("Token not valid")
	}
	return nil
}

func createHash(plainText string) (*string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 14)
	if err != nil {
		return nil, err
	}
	res := string(bytes)
	return &res, err
}

func compareHash(hash, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}

func createToken(userID string) (*UserToken, error) {
	// Refresh Token
	refreshToken := jwt.MapClaims{}
	refreshToken["authorize"] = true
	refreshToken["id"] = userID
	refreshToken["exp"] = time.Now().Add(time.Hour * 24 * 30 * 12 * 5).Unix() // 5 years
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshToken)
	rToken, err := rt.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	// Access Token
	accessToken := jwt.MapClaims{}
	accessToken["authorize"] = true
	accessToken["id"] = userID
	accessToken["exp"] = time.Now().Add(time.Minute * 5).Unix() // 5 Minutes
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessToken)
	aToken, err := at.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &UserToken{
		RefreshToken: rToken,
		AccessToken:  aToken,
	}, nil
}

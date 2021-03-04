package auth

type (
	// Service interface
	Service interface {
		Login(login Login) (*TokenResponse, error)
		Register(register Register) (*TokenResponse, error)
		RefreshToken() (*TokenResponse, error)
		Verify() error
		GetRoles() (*[]UserWithRole, error)
		AddRole(id string) error
		DeleteRole(id string) error
	}
)

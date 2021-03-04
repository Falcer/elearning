package auth

type (
	// Repository interface
	Repository interface {
		GetUserByUsername(username string) (*UserWithPassword, error)
		AddUser(register Register) (*string, error)
		AddUserRole() (*string, error)
		DeleteRoleByID() error
	}
)

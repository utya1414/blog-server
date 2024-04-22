package user

type CreateUser struct {
	Username string
	Email string
	Password string
}

func NewCreateUser(username, email, password string) (*CreateUser, error) {
	return newCreateUser(username, email, password)
}

func newCreateUser(username, email, password string) (*CreateUser, error) {

	if ok, err := IsValidUsername(username); !ok {
		return nil, err
	}

	if ok, err := IsValidEmail(email); !ok {
		return nil, err
	}

	return &CreateUser{
		Username: username,
		Email: email,
		Password: password,
	}, nil
}

func (u *CreateUser) GetUsername() string {
	return u.Username
}

func (u *CreateUser) GetEmail() string {
	return u.Email
}

func (u *CreateUser) GetPassword() string {
	return u.Password
}

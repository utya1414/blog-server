package user

type User struct {
	username string
	email string
	hasshed_password string
	updated_at string
	created_at string
}

func NewUser(username, email, hasshed_password, updated_at, created_at string) (*User, error) {
	return newUser(username, email, hasshed_password, updated_at, created_at)
}

const (
	usernameMinLength = 1
	usernameMaxLength = 20
)

// TODO: 時刻のフォーマットを形式化する
// ビジネスルールによるバリデーションを実装
func newUser(username, email, hasshed_password, updated_at, created_at string) (*User, error) {
	if ok, err := IsValidUsername(username); !ok {
		return nil, err
	}

	if ok, err := IsValidEmail(email); !ok {
		return nil, err
	}

	return &User{
		username: username,
		email: email,
		hasshed_password: hasshed_password,
		updated_at: updated_at,
		created_at: created_at,
	}, nil
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetHasshedPassword() string {
	return u.hasshed_password
}

func (u *User) GetUpdatedAt() string {
	return u.updated_at
}

func (u *User) GetCreatedAt() string {
	return u.created_at
}

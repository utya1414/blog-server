package user

type User struct {
	username string
	email string
	hasshed_password string
}

func NewUser(username, email, hasshed_password string) (*User, error) {
	return newUser(username, email, hasshed_password, )
}

// TODO: 時刻のフォーマットを形式化する
// ビジネスルールによるバリデーションを実装
func newUser(username, email, hasshed_password string) (*User, error) {
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

package user

import (
	"errors"
	"fmt"
	"regexp"
)

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

// ビジネスルールによるバリデーションを実装
func newUser(username, email, hasshed_password, updated_at, created_at string) (*User, error) {
	if ok, err := isValidUsername(username); !ok {
		return nil, err
	}

	if ok, err := isValidEmail(email); !ok {
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

// TODO: バリデーションエラーメッセージの詳細化
func isValidUsername(username string) (bool, error) {
	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return false, fmt.Errorf("ユーザーネームは%d文字以上%d文字以下である必要があります", usernameMinLength, usernameMaxLength)
	}

	// 英数字及びアンダースコアのみ許可
	if !regexp.MustCompile(`^[0-9a-zA-Z_]+$`).MatchString(username) {
		return false, errors.New("ユーザーネームは英数字及びアンダースコアのみ許可されています")
	}
	return true, nil
}

func isValidEmail(email string) (bool, error) {
	// RFC 5322に基づくバリデーション
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return false, errors.New("メールアドレスが不正です")
	}
	return true, nil
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



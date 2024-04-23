package user

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	usernameMinLength = 1
	usernameMaxLength = 20
)

// TODO: バリデーションエラーメッセージの詳細化
func IsValidUsername(username string) (bool, error) {
	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return false, fmt.Errorf("ユーザーネームは%d文字以上%d文字以下である必要があります", usernameMinLength, usernameMaxLength)
	}

	// 英数字及びアンダースコアのみ許可
	if !regexp.MustCompile(`^[0-9a-zA-Z_]+$`).MatchString(username) {
		return false, errors.New("ユーザーネームは英数字及びアンダースコアのみ許可されています")
	}
	return true, nil
}

func IsValidEmail(email string) (bool, error) {
	// RFC 5322に基づくバリデーション
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return false, errors.New("メールアドレスが不正です")
	}
	return true, nil
}
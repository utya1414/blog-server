package user

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/utya1414/blog-server/util"
)

// ビジネスルールによるバリデーションについてテストする責務
func TestNewUser(t *testing.T) {
	type args struct {
		username string
		email string
		hasshed_password string
		updated_at string
		created_at string
	}

	//　ハッシュパスワードについてはテストしない
	hasshed_password := util.RandomPassword()

	testCases := []struct {
		name     string
		args 	args
		want    *User
		// TODO: バリデーションエラーメッセージの詳細化
		wantErr error
	}{
		{
			name: "正常系",
			args: args{
				username: "test_user",
				email: "test@test.com",
				hasshed_password: hasshed_password,
				updated_at: "2021-01-01 00:00:00",
				created_at: "2021-01-01 00:00:00",
			},
			want: &User{
				username: "test_user",
				email: "test@test.com",
				hasshed_password: hasshed_password,
				updated_at: "2021-01-01 00:00:00",
				created_at: "2021-01-01 00:00:00",
			},
			wantErr: nil,
		},
		{
			name: "ユーザーネームの不正",
			args: args{
				username: "",
				email: "test@test.com",
				hasshed_password: hasshed_password,
				updated_at: "2021-01-01 00:00:00",
				created_at: "2021-01-01 00:00:00",
			},
			want: nil,
			wantErr: fmt.Errorf("ユーザーネームは%d文字以上%d文字以下である必要があります", usernameMinLength, usernameMaxLength),
		},
		{
			name: "英数字及びアンダースコア以外のユーザーネーム",
			args: args{
				username: "test_user_test_user_test_user",
				email: "test@test.com",
				hasshed_password: hasshed_password,
				updated_at: "2021-01-01 00:00:00",
				created_at: "2021-01-01 00:00:00",
			},
			want: nil,
			wantErr: fmt.Errorf("ユーザーネームは%d文字以上%d文字以下である必要があります", usernameMinLength, usernameMaxLength),
		},
		{
			name: "ユーザーネームの不正",
			args: args{
				username: "test@user",
				email: "",
				hasshed_password: hasshed_password,
				updated_at: "2021-01-01 00:00:00",
				created_at: "2021-01-01 00:00:00",
			},
			want: nil,
			wantErr: errors.New("ユーザーネームは英数字及びアンダースコアのみ許可されています"),
		},
		{
			name: "イーメールの不正",
			args: args{
				username: "test_user",
				email: "testtest.com",
				hasshed_password: hasshed_password,
				updated_at: "2021-01-01 00:00:00",
				created_at: "2021-01-01 00:00:00",
			},
			want: nil,
			wantErr: fmt.Errorf("メールアドレスが不正です"),
		},
	} 

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := NewUser(testCase.args.username, testCase.args.email, testCase.args.hasshed_password, testCase.args.updated_at, testCase.args.created_at)
			if !reflect.DeepEqual(err, testCase.wantErr) {
				t.Errorf("got error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("got = %v, want %v", got, testCase.want)
			}
		})
	}
}
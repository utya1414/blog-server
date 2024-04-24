package userRep

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	userDomain "github.com/utya1414/blog-server/domain/user"
	"github.com/utya1414/blog-server/util"
)

func TestCreateUserRepository(t *testing.T) {
	okuser, err := userDomain.NewUser(
		util.RandomUsername(),
		util.RandomEmail(),
		util.RandomPassword(),
	)
	require.NoError(t, err)

	tests := []struct {
		name    string
		user    *userDomain.User
		wantErr error
	}{
		{
			name: "正常系",
			user: okuser,
			wantErr: nil,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.CreateUser(ctx, tt.user)
			require.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetUserRepository(t *testing.T) {
	// Todo: 時刻を取得する関数を作成する
	// Todo: 時刻のフォーマットを形式化する
	okuser, err := userDomain.NewUser(
		util.RandomUsername(),
		util.RandomEmail(),
		util.RandomPassword(),
	)
	require.NoError(t, err)
	require.NotNil(t, okuser)

	tests := []struct {
		name    string
		username string
		want   *userDomain.User
		wantErr error
	}{
		{
			name: "正常系",
			username: okuser.GetUsername(),
			want: okuser,
			wantErr: nil,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.CreateUser(ctx, okuser)
			require.NoError(t, err)

			user, err := r.GetUser(ctx, tt.username)
			require.NoError(t, err)
			require.NotNil(t, user)

			require.Equal(t, tt.want.GetUsername(), user.User.GetUsername())
			require.Equal(t, tt.want.GetEmail(), user.User.GetEmail())
			require.Equal(t, tt.want.GetHasshedPassword(), user.User.GetHasshedPassword())
			require.NotNil(t, user.UpdatedAt)
			require.NotNil(t, user.CreatedAt)
		})
	}
}

// func TestListUsersRepository(t *testing.T) {
// 	var okusers []*userDomain.User
// 	for i := 0; i < 5; i++ {
// 		okuser, err := userDomain.NewUser(
// 			util.RandomUsername(),
// 			util.RandomEmail(),
// 			util.RandomPassword(),
// 			"2021-01-01 00:00:00",
// 			"2021-01-01 00:00:00",
// 		)
// 		require.NoError(t, err)
// 		okusers = append(okusers, okuser)
// 	}

// 	// TODO: ユースケース層のパラメータを使うようにする。
// 	arg := *sqlc.ListUsersParams{
// 		Limit: 5,
// 		Offset: 0,
// 	}

// 	tests := []struct {
// 		name    string

		

// }
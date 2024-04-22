package repository

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
		"2021-01-01 00:00:00",
		"2021-01-01 00:00:00",
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
	okuser, err := userDomain.NewUser(
		util.RandomUsername(),
		util.RandomEmail(),
		util.RandomPassword(),
		"2021-01-01 00:00:00",
		"2021-01-01 00:00:00",
	)
	require.NoError(t, err)

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
			require.Equal(t, tt.want.GetUsername(), user.GetUsername())
			require.Equal(t, tt.want.GetEmail(), user.GetEmail())
			require.Equal(t, tt.want.GetHasshedPassword(), user.GetHasshedPassword())
			require.NotNil(t, user.GetUpdatedAt())
			require.NotNil(t, user.GetCreatedAt())
		})
	}
}


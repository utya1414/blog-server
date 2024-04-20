package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/utya1414/blog-server/util"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUsername(),
		Email: util.RandomEmail(),
		HasshedPassword: util.RandomPassword(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err);
	require.NotEmpty(t, user);

	require.Equal(t, arg.Username, user.Username);
	require.Equal(t, arg.Email, user.Email);
	require.Equal(t, arg.HasshedPassword, user.HasshedPassword);

	require.NotNil(t, user.UpdatedAt);
	require.NotNil(t, user.CreatedAt);
	return user;
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t);
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t);
	user2, err := testQueries.GetUser(context.Background(), user1.Username);
	require.NoError(t, err);
	require.NotEmpty(t, user2);

	require.Equal(t, user1.Username, user2.Username);
	require.Equal(t, user1.Email, user2.Email);
	require.Equal(t, user1.HasshedPassword, user2.HasshedPassword);
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second);
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second);
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t);
	}

	arg := ListUsersParams{
		Limit: 5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg);
	require.NoError(t, err);
	require.Len(t, users, 5);

	for _, user := range users {
		require.NotEmpty(t, user.Username);
		require.NotEmpty(t, user.Email);
		require.NotEmpty(t, user.HasshedPassword);
		require.NotNil(t, user.UpdatedAt);
		require.NotEmpty(t, user.CreatedAt);
	}
}
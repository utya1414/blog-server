package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/utya1414/myblog/tree/master/server/util"
)


func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUsername(),
		Email: util.RandomEmail(),
		HasshedPassword: util.RandomPassword(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err);
	return user;
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t);
}
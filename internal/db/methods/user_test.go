package db

import (
	"context"
	"testing"

	"github.com/ChaitanyaSaiV/simple-bank/util"
	"github.com/stretchr/testify/require"
)

// TestCreateDummyAccount creates a dummy account for testing purposes
func createDummyUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	User, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, User)
	require.Equal(t, arg.Username, User.Username)
	require.Equal(t, arg.HashedPassword, User.HashedPassword)
	require.Equal(t, arg.FullName, User.FullName)
	require.Equal(t, arg.Email, User.Email)
	require.NotZero(t, User.CreatedAt)
	return User
}

// TestCreateUser tests the creation of an User
func TestCreateDummyUser(t *testing.T) {
	createDummyUser(t)
}

// TestGetUser tests retrieving an User by ID
func TestGetUser(t *testing.T) {
	User1 := createDummyUser(t)
	User2, err := testQueries.GetUser(context.Background(), User1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, User2)

	require.Equal(t, User1.Username, User2.Username)
	require.Equal(t, User1.HashedPassword, User2.HashedPassword)
	require.Equal(t, User1.FullName, User2.FullName)
	require.Equal(t, User1.Email, User2.Email)
}

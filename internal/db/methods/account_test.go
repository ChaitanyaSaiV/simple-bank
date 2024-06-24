package db

import (
	"context"
	"testing"

	"github.com/ChaitanyaSaiV/simple-bank/util"
	"github.com/stretchr/testify/require"
)

// TestCreateDummyAccount creates a dummy account for testing purposes
func createDummyAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.ID)
	return account
}

// TestCreateAccount tests the creation of an account
func TestCreateAccount(t *testing.T) {
	createDummyAccount(t)
}

// TestGetAccount tests retrieving an account by ID
func TestGetAccount(t *testing.T) {
	account1 := createDummyAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, 0)
}

package db

import (
	"context"
	"testing"

	"github.com/RC-002/Banking_backend/util"
	"github.com/stretchr/testify/require"
)

func RandomTransferHelper(t *testing.T, account1 Account, account2 Account) Transfer {
	id1 := account1.ID
	id2 := account2.ID
	amount := util.RandomAmount()

	args := CreateTransferParams{
		FromAccountID: id1,
		ToAccountID:   id2,
		Amount:        amount,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, id1)
	require.Equal(t, transfer.ToAccountID, id2)
	require.Equal(t, transfer.Amount, amount)

	return transfer
}

func CreateRandomTransfer(t *testing.T) Transfer {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	return RandomTransferHelper(t, account1, account2)

}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
}

func TestGetTransfers(t *testing.T) {
	account1 := CreateRandomAccount(t)

	for i := 0; i < 3; i++ {
		account2 := CreateRandomAccount(t)
		RandomTransferHelper(t, account1, account2)
		RandomTransferHelper(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         3,
		Offset:        2,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 3)

	for _, account := range transfers {
		require.NotEmpty(t, account)
	}
}

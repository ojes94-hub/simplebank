package db

import (
	"context"
	"testing"

	"github.com/ojes94-hub/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, account1.ID)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.Equal(t, transfer.Amount, arg.Amount)

	return transfer

}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	newTransfer := CreateRandomTransfer(t)

	transfer, err := testQueries.GetTransfer(context.Background(), newTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, newTransfer.FromAccountID)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 1; i++ {
		CreateRandomTransfer(t)
	}
	arg := ListTransfersParams{
		Limit:  1,
		Offset: 0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

}

func TestDeleteTransfer(t *testing.T) {
	newTransfer := CreateRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), newTransfer.ID)
	require.NoError(t, err)

}

func TestUpdateTransfer(t *testing.T) {
	newTransfer := CreateRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     newTransfer.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.ID, newTransfer.ID)

}

package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ojes94-hub/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {
	account1 := CreateRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := CreateRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry.Amount, entry2.Amount)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestUpdateEntry(t *testing.T) {
	entry := CreateRandomEntry(t)

	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry2.Amount, arg.Amount)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestListEntries(t *testing.T) {

	arg := ListEntriesParams{
		Limit: 10,
	}

	entry, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

}

func TestDeleteEntry(t *testing.T) {
	entry := CreateRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	entry2, err2 := testQueries.GetEntry(context.Background(), entry.ID)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, entry2)

}

package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/capungkoneng/gomcommerce.git/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAuthor(t *testing.T) Akun {
	arg := CreateAuthorParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	akuntest, err := testQueries.CreateAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, akuntest)
	require.Equal(t, arg.Owner, akuntest.Owner)
	require.Equal(t, arg.Balance, akuntest.Balance)
	require.Equal(t, arg.Currency, akuntest.Currency)

	require.NotZero(t, akuntest.ID)
	require.NotZero(t, akuntest.CreatedAt)

	return akuntest

}

func TestCreateAuthor(t *testing.T) {
	CreateRandomAuthor(t)
}

func TestGetAuthor(t *testing.T) {
	akun1 := CreateRandomAuthor(t)
	akun2, err := testQueries.GetAuthor(context.Background(), akun1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, akun2)

	require.Equal(t, akun1.ID, akun2.ID)
	require.Equal(t, akun1.Owner, akun2.Owner)
	require.Equal(t, akun1.Balance, akun2.Balance)
	require.Equal(t, akun1.Currency, akun2.Currency)
	require.WithinDuration(t, akun1.CreatedAt, akun2.CreatedAt, time.Second)
}

func TestDeleteAuthor(t *testing.T) {
	akun1 := CreateRandomAuthor(t)
	err := testQueries.DeleteAkun(context.Background(), akun1.ID)
	require.NoError(t, err)

	akun2, err := testQueries.GetAuthor(context.Background(), akun1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, akun2)
}

func TestListAuthor(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAuthor(t)
	}

	arg := ListAuthorsParams{
		Limit:  3,
		Offset: 5,
	}

	akun, err := testQueries.ListAuthors(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, akun, 5)

	for _, akun := range akun {
		require.NotEmpty(t, akun)
	}
}

func TestUpdateAuthor(t *testing.T) {
	akun1 := CreateRandomAuthor(t)

	arg := UpdateAkunParams{
		ID:      akun1.ID,
		Balance: util.RandomMoney(),
	}

	akun2, err := testQueries.UpdateAkun(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, akun2)

	require.Equal(t, akun1.ID, akun2.ID)
	require.Equal(t, akun1.Owner, akun2.Owner)
	require.Equal(t, arg.Balance, akun2.Balance)
	require.Equal(t, akun1.Currency, akun2.Currency)
	require.WithinDuration(t, akun1.CreatedAt, akun2.CreatedAt, time.Second)
}

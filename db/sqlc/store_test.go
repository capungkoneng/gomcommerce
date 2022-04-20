package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransfers(t *testing.T) {
	store := NewStore(testDB)

	akun1 := CreateRandomAuthor(t)
	akun2 := CreateRandomAuthor(t)
	fmt.Println(">> before:", akun1.Balance, akun2.Balance)

	//run n
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAkun: akun1.ID,
				ToAkun:   akun2.ID,
				Amount:   amount,
			})
			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, akun1.ID, transfer.FromAkun)
		require.Equal(t, akun2.ID, transfer.ToAkun)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check entry
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, akun1.ID, fromEntry.AkunID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, akun2.ID, toEntry.AkunID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//check akun
		fromAkun := result.FromAkun
		require.NotEmpty(t, fromAkun)
		require.Equal(t, akun1.ID, fromAkun.ID)

		toAkun := result.ToAkun
		require.NotEmpty(t, toAkun)
		require.Equal(t, akun2.ID, toAkun.ID)

		//check akun balance
		fmt.Println(">> tx:", fromAkun.Balance, toAkun.Balance)
		diff1 := akun1.Balance - fromAkun.Balance
		diff2 := toAkun.Balance - akun2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) //1 * amoun, 2 * amount, 3 * amount, ...., n * amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	//check the final updated balances
	updatedAkun1, err := testQueries.GetAuthor(context.Background(), akun1.ID)
	require.NoError(t, err)

	updatedAkun2, err := testQueries.GetAuthor(context.Background(), akun2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAkun1.Balance, updatedAkun2.Balance)
	require.Equal(t, akun1.Balance-int64(n)*amount, updatedAkun1.Balance)
	require.Equal(t, akun2.Balance+int64(n)*amount, updatedAkun2.Balance)

}

func TestTransfersTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	akun1 := CreateRandomAuthor(t)
	akun2 := CreateRandomAuthor(t)
	fmt.Println(">> before:", akun1.Balance, akun2.Balance)

	//run n
	n := 10
	amount := int64(10)

	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAkun := akun1.ID
		toAkun := akun2.ID

		if i%2 == 1 {
			fromAkun = akun2.ID
			toAkun = akun1.ID
		}
		go func() {
			ctx := context.Background()
			_, err := store.TransferTx(ctx, TransferTxParams{
				FromAkun: fromAkun,
				ToAkun:   toAkun,
				Amount:   amount,
			})
			errs <- err
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	//check the final updated balances
	updatedAkun1, err := testQueries.GetAuthor(context.Background(), akun1.ID)
	require.NoError(t, err)

	updatedAkun2, err := testQueries.GetAuthor(context.Background(), akun2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAkun1.Balance, updatedAkun2.Balance)
	require.Equal(t, akun1.Balance, updatedAkun1.Balance)
	require.Equal(t, akun2.Balance, updatedAkun2.Balance)

}

package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore returns a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAkun int64 `json:"from_akun"`
	ToAkun   int64 `json:"to_akun"`
	Amount   int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer  Transfer `json:"transfer"`
	FromAkun  Akun     `json:"from_akun"`
	ToAkun    Akun     `json:"To_akun"`
	FromEntry Entry    `json:"from_entry"`
	ToEntry   Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one akun to the other
// It creates a transfer record, add akun entries, and update akun, balance within a single database transaksi
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfers(ctx, CreateTransfersParams{
			FromAkun: arg.FromAkun,
			ToAkun:   arg.ToAkun,
			Amount:   arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AkunID: arg.FromAkun,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AkunID: arg.ToAkun,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAkun < arg.ToAkun {
			result.FromAkun, result.ToAkun, err = addMoney(ctx, q, arg.FromAkun, -arg.Amount, arg.ToAkun, arg.Amount)

		} else {
			result.ToAkun, result.FromAkun, err = addMoney(ctx, q, arg.ToAkun, arg.Amount, arg.FromAkun, -arg.Amount)

		}

		return nil
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	akunID1 int64,
	amount1 int64,
	akunID2 int64,
	amount2 int64,
) (akun1 Akun, akun2 Akun, err error) {
	akun1, err = q.AddAkunBalance(ctx, AddAkunBalanceParams{
		ID:     akunID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}
	akun2, err = q.AddAkunBalance(ctx, AddAkunBalanceParams{
		ID:     akunID2,
		Amount: amount2,
	})
	return

}

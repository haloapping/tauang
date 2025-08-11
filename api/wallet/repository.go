package wallet

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type walletRepository struct {
	Pool *pgxpool.Pool
}

func NewWalletRepository(pool *pgxpool.Pool) walletRepository {
	return walletRepository{
		Pool: pool,
	}
}

type walletDB struct {
	ID          string
	UserID      string
	Name        string
	Description string
	Currency    string
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

func (r walletRepository) create(c echo.Context, req createWalletReq) (walletDB, error) {
	tx, err := r.Pool.BeginTx(c.Request().Context(), pgx.TxOptions{})
	if err != nil {
		return walletDB{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(c.Request().Context())
		} else {
			tx.Commit(c.Request().Context())
		}
	}()

	ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
	ib.InsertInto("wallets")
	ib.Cols("id", "user_id", "name", "description", "currency")
	ID := uuid.NewString()
	ib.Values(ID, req.UserID, req.Name, req.Description, req.Currency)
	ib.Returning("*")
	q, args := ib.Build()

	var w walletDB
	row := tx.QueryRow(c.Request().Context(), q, args...)
	err = row.Scan(&w.ID, &w.UserID, &w.Name, &w.Description, &w.Currency, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		return walletDB{}, err
	}

	return w, nil
}

func (r walletRepository) findByID(c echo.Context, ID string) (walletDB, error) {
	tx, err := r.Pool.BeginTx(c.Request().Context(), pgx.TxOptions{})
	if err != nil {
		return walletDB{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(c.Request().Context())
		} else {
			tx.Commit(c.Request().Context())
		}
	}()

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "user_id", "name", "description", "currency", "created_at", "updated_at")
	sb.From("wallets")
	sb.Where(sb.Equal("id", ID))
	q, args := sb.Build()

	row := tx.QueryRow(c.Request().Context(), q, args...)
	var w walletDB
	err = row.Scan(&w.ID, &w.UserID, &w.Name, &w.Description, &w.Currency, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		return walletDB{}, err
	}

	return w, nil
}

func (r walletRepository) findAll(c echo.Context, req createWalletReq) ([]walletDB, error) {
	tx, err := r.Pool.BeginTx(c.Request().Context(), pgx.TxOptions{})
	if err != nil {
		return []walletDB{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(c.Request().Context())
		} else {
			tx.Commit(c.Request().Context())
		}
	}()

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "user_id", "name", "description", "currency", "created_at", "updated_at")
	sb.From("wallets")
	q, args := sb.Build()

	rows, err := tx.Query(c.Request().Context(), q, args...)
	if err != nil {
		return []walletDB{}, err
	}
	var wallets []walletDB
	for rows.Next() {
		var w walletDB
		err := rows.Scan(&w.ID, &w.UserID, &w.Name, &w.Description, &w.Currency, &w.CreatedAt, &w.UpdatedAt)
		if err != nil {
			return []walletDB{}, err
		}
		wallets = append(wallets, w)
	}

	return wallets, nil
}

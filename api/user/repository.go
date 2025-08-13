package user

import (
	"github.com/google/uuid"
	"github.com/guregu/null/v6"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	Pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) userRepository {
	return userRepository{
		Pool: pool,
	}
}

type userRegisterDB struct {
	ID        string
	Username  string
	Email     string
	Phone     string
	Password  string
	CreatedAt null.Time
	UpdatedAt null.Time
}

func (r userRepository) register(c echo.Context, req registerReq) (userRegisterDB, error) {
	tx, err := r.Pool.BeginTx(c.Request().Context(), pgx.TxOptions{})
	if err != nil {
		return userRegisterDB{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(c.Request().Context())
		} else {
			tx.Commit(c.Request().Context())
		}
	}()

	// insert new user
	ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
	ib.InsertInto("users")
	ib.Cols("id", "username", "email", "phone", "password")
	hashPasswordByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return userRegisterDB{}, err
	}
	ID := uuid.NewString()
	ib.Values(ID, req.Username, req.Email, req.Phone, string(hashPasswordByte))
	q, args := ib.Build()

	_, err = tx.Exec(c.Request().Context(), q, args...)
	if err != nil {
		return userRegisterDB{}, err
	}

	// select new user
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "username", "email", "phone", "password", "created_at", "updated_at")
	sb.From("users")
	sb.Where(sb.Equal("id", ID))
	q, args = sb.Build()
	row := tx.QueryRow(c.Request().Context(), q, args...)
	var u userRegisterDB
	err = row.Scan(&u.ID, &u.Username, &u.Email, &u.Phone, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return userRegisterDB{}, err
	}

	return u, nil
}

type userLoginDB struct {
	Username string
	Password string
}

func (r userRepository) login(c echo.Context, req loginReq) (userLoginDB, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("username", "password")
	sb.From("users")
	sb.Where(sb.Equal("username", req.Username))
	q, args := sb.Build()

	tx, err := r.Pool.BeginTx(c.Request().Context(), pgx.TxOptions{})
	if err != nil {
		return userLoginDB{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(c.Request().Context())
		} else {
			tx.Commit(c.Request().Context())
		}
	}()

	row := tx.QueryRow(c.Request().Context(), q, args...)
	var u userLoginDB
	err = row.Scan(&u.Username, &u.Password)
	if err != nil {
		return userLoginDB{}, err
	}

	tx.Commit(c.Request().Context())

	return u, nil
}

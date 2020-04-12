package account

import (
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

type Account struct {
	ID        string
	Fullname  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostgresDB struct {
	DB *sql.DB
}

func (postgres PostgresDB) SignUp(newAccount Account, now time.Time) (Account, error) {
	account := Account{
		Fullname:  newAccount.Fullname,
		Email:     newAccount.Email,
		Password:  newAccount.Password,
		CreatedAt: now.UTC(),
		UpdatedAt: now.UTC(),
	}

	query := `INSERT INTO account(Fullname, Email, Password) VALUES($1, $2, $3) RETURNING ID, Fullname, Email, Password, CreatedAt, UpdatedAt`

	err := postgres.DB.QueryRow(query, account.Fullname, account.Email, account.Password).Scan(&account)

	return account, err
}

func (postgres PostgresDB) GetProductByID(id string) ([]Account, error) {
	var accounts []Account

	query := `SELECT * FROM account`

	rows, err := postgres.DB.Query(query)

	rows.Scan(&accounts)
	if err != nil {
		panic(err)
	}

	return accounts, err
}

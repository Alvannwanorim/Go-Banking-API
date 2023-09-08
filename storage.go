package main

import (
	"database/sql"
	"fmt"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := ` CREATE TABLE  IF NOT EXISTS account (
		id serial PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		account_number serial ,
		balance INTEGER NOT NULL DEFAULT 0,
		created_at TIMESTAMP DEFAULT NOW()
	)
	`
	_, err := s.db.Exec(query)
	return err
}

func NewPostgresStore() (*PostgresStore, error) {
	consStr := "user=postgres dbname=go-bank password=alvan2327 sslmode=disable"
	db, err := sql.Open("postgres", consStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}
func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := (`
		INSERT INTO account(first_name, last_name, account_number, balance)
		VALUES($1,$2,$3,$4)
	`)

	resp, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.AccountNumber, acc.Balance)

	if err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query(`DELETE FROM account WHERE id=$1`, id)

	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query(`
		SELECT * FROM account WHERE id=$1
	`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return ScanAccount(rows)
	}
	return nil, fmt.Errorf("account with id: %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query(`SELECT * FROM account`)

	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account, err := ScanAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func ScanAccount(row *sql.Rows) (*Account, error) {
	account := new(Account)

	err := row.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.AccountNumber,
		&account.Balance,
		&account.CreatedAt,
	)

	return account, err
}

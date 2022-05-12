package model

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) CreateUser(u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO user VALUES(null, ?, ?);`
	_, err := m.DB.ExecContext(ctx, query, u.Username, u.Password)

	if err != nil {
		return err
	}
	return nil
} 

func (m *DBModel) GetUser(username string) (*UserInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, username FROM user WHERE username = ?`
	row := m.DB.QueryRowContext(ctx, query, username)

	var user UserInfo
	err := row.Scan(
		&user.Id,
		&user.Username,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
} 

func (m *DBModel) GetUserWithPassword(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	
	query := `SELECT id, username, password FROM user WHERE username = ?`
	row := m.DB.QueryRowContext(ctx, query, username)

	var user User
	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
} 
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

func (m *DBModel) GetNotifications(uid int) ([]*Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, name, user_id FROM notification WHERE user_id = ?`
	rows, err := m.DB.QueryContext(ctx, query, uid)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var notifications []*Notification
	for rows.Next() {
		var notification Notification
		err = rows.Scan(
			&notification.Id,
			&notification.Title,
			&notification.Name,
			&notification.UserId,
		)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	if err != nil {
		return nil, err
	}
	return notifications, nil
} 
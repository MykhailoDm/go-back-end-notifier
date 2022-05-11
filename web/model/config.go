package model

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Port    string
	Env     string
	Version string
	JwtExpirationTime int
	JwtKey string
	db struct {
		dsn string
		name string
	}
}

func (c *Config) GetConfig() {
	c.Version = readEnv("Version", "N/A")
	c.Port = readEnv("Port", "4000")
	c.Env = readEnv("ENV", "dev")
	c.JwtExpirationTime, _ = strconv.Atoi(readEnv("EXPIRATION_DATE", "1440"))
	c.JwtKey = readEnv("JWT_KEY", "test_key")
	c.db.dsn = readEnv("DSN", "root:Gek313nkL@tcp(127.0.0.1:3306)/go_sql_notifier")
	c.db.name = readEnv("DB_NAME", "mysql")
}

func readEnv(varName string, dft string) string {
	v := os.Getenv(varName)
	if v == "" {
		return dft
	}
	return varName
}



func OpenDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.db.name, cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
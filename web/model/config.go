package model

import (
	"os"
	"strconv"
)

type Config struct {
	Port    string
	Env     string
	Version string
	JwtExpirationTime int
	JwtKey string
}

func (c *Config) GetConfig() {
	c.Version = readEnv("Version", "N/A")
	c.Port = readEnv("Port", "4000")
	c.Env = readEnv("Env", "dev")
	c.JwtExpirationTime, _ = strconv.Atoi(readEnv("EXPIRATION_DATE", "1440"))
	c.JwtKey = readEnv("JWT_KEY", "test_key")
}

func readEnv(varName string, dft string) string {
	v := os.Getenv(varName)
	if v == "" {
		return dft
	}
	return varName
}
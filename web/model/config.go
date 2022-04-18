package model

import "os"

type Config struct {
	Port    string
	Env     string
	Version string
}

func (c *Config) GetConfig() {
	c.Version = readEnv("Version", "N/A")
	c.Port = readEnv("Port", "4000")
	c.Env = readEnv("Env", "dev")
}

func readEnv(varName string, dft string) string {
	v := os.Getenv(varName)
	if v == "" {
		return dft
	}
	return varName
}
package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	MySqlDSN string
}

func InitConfig() *Config {
	return &Config{
		MySqlDSN: GetEnv("MYSQL_DSN", "root:secret@tcp(localhost:3306)/simple-golang-echo?parseTime=true"),
	}
}

func GetEnv(key, defaultValue string) string {
	env := os.Getenv(key)

	if env != "" {
		return env
	}

	return defaultValue
}

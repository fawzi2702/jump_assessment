package environement

import (
	"fmt"
	"os"
	"strconv"
)

type env struct {
	Mode string

	// DB
	DB_Name     string
	DB_User     string
	DB_Password string
	DB_Host     string
	DB_Port     string

	API_Port string
}

var environment *env
var environmentSetup = false

func (e *env) validate() error {
	if e.Mode != "dev" && e.Mode != "prod" {
		return fmt.Errorf("MODE environement variable must be either dev or prod")
	}

	if e.DB_Name == "" {
		return fmt.Errorf("DB_Name environement variable must be provided")
	}
	if e.DB_User == "" {
		return fmt.Errorf("DB_User environement variable must be provided")
	}
	if e.DB_Password == "" {
		return fmt.Errorf("DB_Password environement variable must be provided")
	}
	if e.DB_Host == "" {
		return fmt.Errorf("DB_Host environement variable must be provided")
	}
	if e.DB_Port == "" {
		return fmt.Errorf("DB_PORT environement variable must be provided")
	}

	if e.API_Port == "" {
		return fmt.Errorf("API_PORT environement variable must be provided")
	} else if p, err := strconv.Atoi(e.API_Port); err != nil || p < 0 || p > 65535 {
		return fmt.Errorf("API_PORT environement variable must be a valid port number")
	}

	return nil
}

func (e *env) dbDsn() (string, error) {
	if !environmentSetup {
		return "", fmt.Errorf("environment not loaded")
	}

	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", e.DB_Host, e.DB_User, e.DB_Password, e.DB_Name, e.DB_Port)

	return uri, nil
}

func LoadEnv() error {
	if environmentSetup {
		return nil
	}

	// if err := godotenv.Load(); err != nil {
	// 	return err
	// }

	environment = &env{
		Mode:        os.Getenv("MODE"),
		DB_Name:     os.Getenv("POSTGRES_DB"),
		DB_User:     os.Getenv("POSTGRES_USER"),
		DB_Password: os.Getenv("POSTGRES_PASSWORD"),
		DB_Host:     os.Getenv("POSTGRES_HOST"),
		DB_Port:     os.Getenv("POSTGRES_PORT"),
		API_Port:    os.Getenv("API_PORT"),
	}

	if err := environment.validate(); err != nil {
		return err
	}

	environmentSetup = true

	return nil
}

func Get(key string) (string, error) {
	if !environmentSetup {
		return "", fmt.Errorf("environment not loaded")
	}

	switch key {
	case "MODE":
		return environment.Mode, nil
	case "DB_DSN":
		return environment.dbDsn()
	case "API_PORT":
		return environment.API_Port, nil
	default:
		return "", fmt.Errorf("key <%s> not found", key)
	}
}

func GetInt(key string) (int, error) {
	if !environmentSetup {
		return 0, fmt.Errorf("environment not setup")
	}

	strValue, err := Get(key)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, fmt.Errorf("key <%s> is not a valid integer", key)
	}

	return value, nil
}

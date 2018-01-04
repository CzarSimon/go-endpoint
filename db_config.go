package endpoint

import (
	"database/sql"
	"fmt"
)

const (
	DB_USER_KEY = "USER"
	DB_PWD_KEY  = "PASSWORD"
	DB_NAME_KEY = "NAME"

	DEFAULT_POSTGRES_PORT = "5432"
	POSTGRES_PROTOCOL     = "tcp"
)

// DBConfig Interface for database configuration and connection
type DBConfig interface {
	Connect() (*sql.DB, error)
}

// PGConfig Connection info for a postgres database
type PGConfig struct {
	ServerAddr
	User     string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// NewPGConfig
func NewPGConfig(name string) PGConfig {
	host := NewServerAddr(name)
	host.Protocol = POSTGRES_PROTOCOL
	if host.Port == "" {
		host.Port = DEFAULT_POSTGRES_PORT
	}
	return PGConfig{
		ServerAddr: host,
		User:       getenv(makeKey(name, DB_USER_KEY), ""),
		Password:   getenv(makeKey(name, DB_PWD_KEY), ""),
		Database:   getenv(makeKey(name, DB_NAME_KEY), "postgres"),
	}
}

// Connect Attempts to connect to a postgres database. Requires a driver!
func (pg PGConfig) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		pg.Host, pg.User, pg.Password, pg.Database, pg.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	return db, err
}

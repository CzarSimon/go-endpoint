package endpoint

import (
	"database/sql"
	"fmt"
	"os"
)

// Database environment variable keyssuffixes
const (
	DB_USER_KEY     = "USER"
	DB_PWD_KEY      = "PASSWORD"
	DB_NAME_KEY     = "NAME"
	DB_SSL_MODE_KEY = "SSL_MODE"
	DB_VERSION      = "VERSION"
)

// PostgreSQL default values
const (
	DEFAULT_POSTGRES_PORT     = "5432"
	DEFAULT_POSTGRES_SSL_MODE = "disable"
	DEFAULT_POSTGRES_DATABASE = "postgres"
	POSTGRES_PROTOCOL         = "tcp"
)

// MySQL default values
const (
  	DEFAULT_MYSQL_PORT     = "3306"
  	DEFAULT_MYSQL_PROTOCOL = "tcp"
)

// SQLite default values
const (
	DEFAULT_SQLITE_DRIVER = "sqlite3"
)

// ConnInfo Information needed to connect to a database
type ConnInfo struct {
	DriverName string `json:"driverName"`
	ConnStr    string `json:"connStr"`
}

// SQLConfig Interface for database configuration and connection
type SQLConfig interface {
	Connect() (*sql.DB, error)
	ConnInfo() ConnInfo
}

// PGConfig Connection info for a postgres database
type PGConfig struct {
	ServerAddr
	User     string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslMode"`
}

// NewPGConfig Creates a new PGConfig by reading from environment variables
func NewPGConfig(name string) PGConfig {
	host := NewServerAddr(name)
	host.Protocol = POSTGRES_PROTOCOL
	if host.Port == "" {
		host.Port = DEFAULT_POSTGRES_PORT
	}
	return PGConfig{
		ServerAddr: host,
		User:       os.Getenv(makeKey(name, DB_USER_KEY)),
		Password:   os.Getenv(makeKey(name, DB_PWD_KEY)),
		Database:   getenv(makeKey(name, DB_NAME_KEY), DEFAULT_POSTGRES_DATABASE),
		SSLMode:    getenv(makeKey(name, DB_SSL_MODE_KEY), DEFAULT_POSTGRES_SSL_MODE),
	}
}

// ConnInfo Creates and returns connection info from a PGConfig
func (pg PGConfig) ConnInfo() ConnInfo {
	return ConnInfo{
		DriverName: "postgres",
		ConnStr:    pg.getDSN(),
	}
}

// getDSN Structures and returns a PostgreSQL datasource name. Only TCP connection is supported
func (pg PGConfig) getDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		pg.Host, pg.User, pg.Password, pg.Database, pg.Port, pg.SSLMode)
}

// Connect Attempts to connect to a postgres database. Requires a driver!
func (pg PGConfig) Connect() (*sql.DB, error) {
	return connectDB(pg.ConnInfo())
}

// SQLiteConfig Configuration info for a SQLite db
type SQLiteConfig struct {
	DriverVersion string `json:"driverVersion"`
	File string `json:"file"`
}

// NewSQLiteConfig Creates a new SQLconfig by reading from environment variables
func NewSQLiteConfig(name string) SQLiteConfig {
	return SQLiteConfig{
		Version: getenv(makeKey(name, DB_VERSION), DEFAULT_SQLITE_DRIVER),
		File: 	 os.Getenv(makeKey(name, DB_NAME)),
	}
}

// Connect Connects to a sqlite database. Requires a driver!
func (lite SQLiteConfig) Connect() (*sql.DB, error) {
	return connectDB(lite.ConnInfo())
}

// ConnInfo Gets the connection information 
func (lite SQLiteConfig) ConnInfo() ConnInfo {
	return ConnInfo{
		DriverName: lite.DriverVersion,
		ConnStr:    lite.File,
	}
}

// MySQLConfig Configuration info for a MySQL db
type MySQLConfig struct {
  ServerAddr
  User     string `json:"username"`
  Password string `json:"password"`
  Database string `json:"database"`
}

// NewMySQLConfig Creates a new MySQLConfig from environment variables
func NewMySQLConfig(name string) MySQLConfig {
	server := NewServerAddr(name)
	server.Protocol = DEFAULT_MYSQL_PROTOCOL
	return MySQLConfig{
		ServerAddr: server,
		User:       os.Getenv(makeKey(name, DB_USER_KEY)),
		Password:   os.Getenv(makeKey(name, DB_PWD_KEY)),
		Database:   os.Getenv(makeKey(name, DB_NAME_KEY)),
	}
}

// Connect Attempts to connect to a MySQL database. Requires a driver!
func (my MySQLConfig) Connect() (*sql.DB, error) {
  return connectDB(my.ConnInfo())
}

// ConnInfo Gets the connection information
func (my MySQLConfig) ConnInfo() ConnInfo {
  return ConnInfo{
    DriverName: "mysql",
    ConnStr:    my.getDSN(),
  }
}

// getDSN Structures and returns a MySQL datasource name. Only TCP connection is supported
func (my MySQLConfig) getDSN() string {
  return fmt.Sprintf("%s:%s@%s(%s%s)/%s",
    my.User, my.Password, my.Protocol, my.Host, my.getPortString() my.Database)
}

// connectDB Generic connection and ping test method for a SQL db
func connectDB(conn ConnInfo) (*sql.DB, error) {
	db, err := sql.Open(conn.DriverName, conn.ConnStr)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	return db, err
}

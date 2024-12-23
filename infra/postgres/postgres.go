package infra

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	ErrMigrationFailed = errors.New("database migration failed")
)

type Postgre struct {
	Username string
	Password string
	Port     string
	Address  string
	Database string

	DB *sqlx.DB
}

type PsqlDb struct {
	*Postgre
}

var PSQL *PsqlDb

func InitPostgre() error {
	PSQL = new(PsqlDb)

	PSQL.Postgre = &Postgre{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	if err := PSQL.Postgre.OpenConnection(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// if err := PSQL.runMigration("./migration/migration.sql"); err != nil {
	// 	return fmt.Errorf("%w: %s", ErrMigrationFailed, err)
	// }

	log.Println("Database initialized successfully")
	return nil
}

func (p *Postgre) OpenConnection() error {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		p.Username, p.Password, p.Address, p.Port, p.Database,
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	p.DB = db
	return nil
}

func (p *PsqlDb) runMigration(filePath string) error {
	query, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	_, err = p.DB.Exec(string(query))
	if err != nil {
		return fmt.Errorf("failed to execute migration query: %w", err)
	}

	return nil
}

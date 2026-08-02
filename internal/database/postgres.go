package database

import (
	"database/sql"
	"fmt"

	"github.com/DmitrySychevDev/career-tracker/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB    *gorm.DB
	SQLDB *sql.DB
}

func NewPostgres(cfg *config.Config) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.DatabaseHost,
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseName,
		cfg.DatabasePort,
		cfg.DatabaseSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{
		DB:    db,
		SQLDB: sqlDB,
	}, nil
}

func (p *Postgres) Close() error {
	return p.SQLDB.Close()
}

package postgres

import (
	"SmartHouseAPI/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(config config.DBConfig) (*Repository, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password =%s dbname=%s sslmode=disable", config.Host, config.Port, config.Username, config.Password, config.Database)
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &Repository{DB: db}, nil
}

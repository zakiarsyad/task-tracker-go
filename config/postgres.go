package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct{}

func NewDB() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Connect() (*gorm.DB, error) {
	dsn := "postgres://zakiarsyad:postgres@localhost:5432/neurons"

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

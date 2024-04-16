package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/Aerok925/LeoNews/configs"
)

func NewDB(cfg configs.DB) (*sql.DB, error) {
	dst := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DB,
		cfg.SSL)
	db, err := sql.Open("postgres", dst)
	if err != nil {
		return nil, err
	}
	return db, nil
}

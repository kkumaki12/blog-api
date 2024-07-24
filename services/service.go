package services

import "database/sql"

type MyAppService struct {
	db *sql.DB
}

// コンストラクタ
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}

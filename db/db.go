package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wanta-zulfikri/Projek_unit1/config"
	"github.com/wanta-zulfikri/Projek_unit1/helper"
)

func InitDb() *sql.DB {
	db, err := sql.Open(config.DbDriver, fmt.Sprintf("root:%s@tcp(localhost:%s)/%s?parseTime=true", config.DbPassword, config.DbPort, config.DBname))
	helper.PanicIfError(err)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}

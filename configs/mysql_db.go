package configs

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL(mySqlDSN string) *sql.DB {
	db, err := sql.Open("mysql", mySqlDSN)
	if err != nil {
		for i := 1; i <= 20; i++ {
			fmt.Printf("trying to reconnect database #%d\n", i)
			db, err = sql.Open("mysql", mySqlDSN)
			if err == nil {
				break
			}

			time.Sleep(5 * time.Second)
		}

		if err != nil {
			panic(err)
		}
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

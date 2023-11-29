package migrations

import (
	"context"
	"database/sql"
	"fmt"
)

func MigrateTable(ctx context.Context, db *sql.DB) {
	if _, err := db.ExecContext(ctx, books); err != nil {
		panic(err)
	}

	fmt.Println("migrate table success")
}

var books = `
CREATE TABLE IF NOT EXISTS books (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)
`

package sqlx

import (
	"database/sql"
	"time"
)

func OpenDB(driver, databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open(driver, databaseUrl)
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(2)
	db.SetConnMaxLifetime(5*time.Minute)
	return db, nil
}

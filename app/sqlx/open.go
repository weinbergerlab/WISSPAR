package sqlx

import (
	"entgo.io/ent/dialect/sql"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
	"time"
)

func Open(driver, databaseUrl string) (*models.Client, error) {
	drv, err := sql.Open(driver, databaseUrl)
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(2)
	db.SetConnMaxLifetime(5*time.Minute)
	return models.NewClient(models.Driver(drv)), nil
}


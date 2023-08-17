package migration

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitMigrations() error {
	db, err := sql.Open(
		viper.GetString("db.driver"),
		fmt.Sprintf("%s://%s:%s@%s",
			viper.GetString("db.type"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.dataSource")),
	)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	dbInstance, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", viper.GetString("db.filePath")),
		viper.GetString("db.driver"),
		driver)
	if err != nil {
		return err
	}

	if err := dbInstance.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

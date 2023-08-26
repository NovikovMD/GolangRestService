package repository

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"os"
)

func initMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	dbInstance, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", viper.GetString("db.filePath")),
		"postgres",
		driver)
	if err != nil {
		return err
	}

	if err = dbInstance.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func NewPostgresInstance() (*sqlx.DB, error) {
	dbInstance, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			viper.GetString("db.host"),
			viper.GetString("db.port"),
			viper.GetString("db.user"),
			viper.GetString("db.dbname"),
			os.Getenv("DB_PASSWORD"),
			viper.GetString("db.sslmode")))
	if err != nil {
		return nil, err
	}

	err = dbInstance.Ping()
	if err != nil {
		return nil, err
	}

	if err = initMigrations(dbInstance); err != nil {
		return nil, err
	}

	return dbInstance, nil
}

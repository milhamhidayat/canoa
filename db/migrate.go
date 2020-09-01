package db

import (
	"database/sql"
	"errors"
	"path"
	"runtime"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/golang-migrate/migrate/v4"
	_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Get file migration
)

type migration struct {
	Migrate *migrate.Migrate
}

// Up is function to run database migrations
func (m *migration) Up() (bool, error) {
	err := m.Migrate.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return true, nil
		}
		return false, err
	}
	return true, nil
}

// Down is function for migrate database to previous version
func (m *migration) Down() (bool, error) {
	err := m.Migrate.Down()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *migration) Drop() error {
	return m.Migrate.Drop()
}

func runMigration(dbConn *sql.DB, migrationsFolder string) (*migration, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("fail to run caller")
	}

	migrationPath := path.Join(path.Dir(filename), migrationsFolder)

	driver, err := _mysql.WithInstance(dbConn, &_mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationPath, "mysql", driver)
	if err != nil {
		return nil, err
	}
	return &migration{Migrate: m}, nil
}

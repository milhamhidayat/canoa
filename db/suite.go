package db

import (
	"database/sql"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Suite is struct for mariadb test suite
type Suite struct {
	suite.Suite
	DBConn    *sql.DB
	DBName    string
	Migration *migration
}

// SetupSuite will initialize test suite
func (s *Suite) SetupSuite() {
	dsn := "soccer:soccer-pass@tcp(localhost:3306)/soccer?parseTime=1&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	migrationsFolder := "migrations"

	dbConn, err := sql.Open("mysql", dsn)
	require.NoError(s.T(), err)

	s.DBConn = dbConn
	err = s.DBConn.Ping()
	require.NoError(s.T(), err)

	s.Migration, err = runMigration(s.DBConn, migrationsFolder)
	require.NoError(s.T(), err)
}

// TearDownSuite will close db connection
func (s *Suite) TearDownSuite() {
	err := s.Migration.Migrate.Drop()
	require.NoError(s.T(), err)

	err = s.DBConn.Close()
	require.NoError(s.T(), err)
}

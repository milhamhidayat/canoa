package db

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

// Open is createing mysql connection
func Open(dsn string) *sql.DB {
	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("connect to MySQL")
	}

	err = dbConn.Ping()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("ping MySQL")
	}
	log.Info("connected to MySQL")

	return dbConn
}

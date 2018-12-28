package database

import (
	"database/sql"
	"github.com/globalsign/mgo"
)

var (
	DBCon *sql.DB
	MBCon *mgo.Database
)
package framework

import (
	"github.com/ivanj4u/service-switch/database"
)

func Init() {
	// Import Properties
	initProperties()

	// Open Database
	database.DBCon = openDatabaseConnection()
	database.MBCon = openMongoDBConnection()

	// Loader
	load()

	// Mapping URL
	mappingUrl()
}
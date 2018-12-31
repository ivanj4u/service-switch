package framework

func Init() {
	// Import Properties
	initProperties()

	// Open Database Connection
	DBCon = openDatabaseConnection()

	// Open MongoDB Connection
	initMongoDBConnection()
	defer closeMongoDBConnection()

	// Loader
	load()

	// Mapping URL
	mappingUrl()
}
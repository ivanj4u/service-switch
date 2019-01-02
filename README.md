# Switch Service Application

This application used for routing all request to the core services


## Built With

* [go](https://golang.org/doc/) - Core Framework
* [go-sql-driver](https://github.com/go-sql-driver/mysql) - MySQL Driver
* [go-mongodb](https://github.com/globalsign/mgo) - MongoDB Driver
* [natefinch-lumberjack](https://github.com/natefinch/lumberjack) - Logger


### Installing

Make sure you already install go and your **$GOPATH** is right.

```
go env
```

Get and install MySQL Driver
```
go get github.com/go-sql-driver/mysql
cd $GOPATH/src/github.com/go-sql-driver/mysql
go install
```

Get and install MongoDB Driver
```
go get github.com/globalsign/mgo
cd $GOPATH/src/github.com/globalsign/mgo
go install
```

Get and install Logger
```
go get github.com/natefinch/lumberjack
cd $GOPATH/src/github.com/natefinch/lumberjack
go install
```

### Build
```
cd $GOPATH/src/github.com/ivanj4u/service-switch
go install
go build
```
### Configuration
File configuration
```
cd $GOPATH/build/cfg
```
Here some example of configuration
```
## Application Properties
application.port=9090

## Logger
app.log=log/application.log
app.size=1024

# Database Default
db.url=127.0.0.1
db.port=3306
db.driver=mysql
db.dbname=test-db
db.user=user
db.pass=pass
db.schema=TEST

#Database Logging
mongo.url=127.0.0.1
mongo.port=27017
mongo.dbname=test-log
mongo.user=user
mongo.pass=pass

```

### Running

Go to folder build
```
cd $GOPATH/build
```
Windows
```
app.exe
```
Linux
```
sh app
```

## Authors

* **Ivan Aribanilia** - *Developer* - [Githubs](https://github.com/ivanj4u)

## License

Copyright 2018 [Ivan Aribanilia](mailto:angko.j4u@gmail.com).
Please tell me if you interested with this project

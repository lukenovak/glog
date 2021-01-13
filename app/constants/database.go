package constants

import (
	"fmt"
	"os"
)

const POSTGRES = "postgres"

var dbHost = os.Getenv("GLOG_DB_HOST")
var dbPort = os.Getenv("GLOG_DB_PORT")
var dbUser = os.Getenv("GLOG_DB_USER")
var dbPass = os.Getenv("GLOG_DB_PASS")
var dbName = os.Getenv("GLOG_DB_NAME")

var PsqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	dbHost, dbPort, dbUser, dbPass, dbName)
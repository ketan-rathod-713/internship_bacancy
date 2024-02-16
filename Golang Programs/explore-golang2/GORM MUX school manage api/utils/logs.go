package utils

import (
	"log"
	"os"
)

var ApiLogs = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var ErrorLogs = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

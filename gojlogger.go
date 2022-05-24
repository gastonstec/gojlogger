// Package gojlogger provides a custom logger with
// Info, Warning an Error logging methods.
//
// The default timezone is UTC, use
// TimeZoneUTC = false for system timezone.
//
// Package usage:
//   1. Initialize logger with InitLogger function
//   2. Use LogInfo, LogWarning and LogError with
//      a plain or JSON string
package gojlogger

import (
	"log"
	"os"
	"time"
	"fmt"
	"github.com/gastonstec/goutils"
)

var (
	cLogger     	*log.Logger = nil
	TimeZoneUTC 	bool        = true
)

const(
	CANNOT_OPEN_FILE = "%s: cannot open file %s"
)

// Function InitLogger init the logger with the specified path
// and filename (for example ./mylog.log), if the path is empty
// os.Stdout will be used. The log file is open for append.
func InitLogger(filepath string) error {

	// check filepath
	if filepath == "" {
		// log to stdout
		cLogger = log.New(os.Stdout, "", 0)
	} else {
		// open log file for append
		logFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return fmt.Errorf(CANNOT_OPEN_FILE, goutils.GetFunctionName(), filepath)
		}
		cLogger = log.New(logFile, "", 0)
	}

	// return no errors
	return nil
}

// Function LogInfo writes a INFO entry in the log
func LogInfo(msg string) {

	var logEntry string
	var entryTime string

	// check for timezone an assign value
	if TimeZoneUTC {
		entryTime = time.Now().UTC().String()
	} else {
		entryTime = time.Now().String()
	}

	// check for JSON or plain string
	if goutils.IsJSON(msg) {
		logEntry = "{\"datetime\": \"" + entryTime + "\", \"level\": \"INFO\", \"message\": " + msg + "}"
	} else {
		logEntry = "{\"datetime\": \"" + entryTime + "\", \"level\": \"INFO\", \"message\": \"" + msg + "\"}"
	}

	// insert log entry
	cLogger.Println(logEntry)

}

// Function LogInfo writes a WARNING entry in the log
func LogWarning(msg string) {

	var logEntry string
	var entryTime string

	// check for timezone an assign value
	if TimeZoneUTC {
		entryTime = time.Now().UTC().String()
	} else {
		entryTime = time.Now().String()
	}

	// check for JSON or plain string
	if goutils.IsJSON(msg) {
		logEntry = "{\"datetime\": \"" + entryTime + "\", \"level\": \"WARNING\", \"message\": " + msg + "}"
	} else {
		logEntry = "{\"datetime\": \"" + entryTime + "\", \"level\": \"WARNING\", \"message\": \"" + msg + "\"}"
	}

	// insert log entry
	cLogger.Println(logEntry)

}

// Function LogInfo writes an ERROR entry in the log
func LogError(msg string) {

	var logEntry string
	var entryTime string

	// check for timezone an assign value
	if TimeZoneUTC {
		entryTime = time.Now().UTC().String()
	} else {
		entryTime = time.Now().String()
	}

	// check for JSON or plain string
	if goutils.IsJSON(msg) {
		logEntry = "{\"datetime\": \"" + entryTime + "\", \"level\": \"ERROR\", \"message\": " + msg + "}"
	} else {
		logEntry = "{\"datetime\": \"" + entryTime + "\", \"level\": \"ERROR\", \"message\": \"" + msg + "\"}"
	}

	// insert log entry
	cLogger.Println(logEntry)

}

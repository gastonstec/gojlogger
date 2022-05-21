// Copyright Kueski. All rights reserved.
// Use of this source code is not licensed

// Package clogger provides a custom logger with
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
	"github.com/gastonstec/goutils"
)

var (
	cLogger     *log.Logger = nil
	TimeZoneUTC bool        = true
)

// Function InitLogger init the logger with the specified path
// and filename (for example ./mylog.log), if the path is empty
// os.Stdout will be used
func InitLogger(filepath string) error {

	//Check filepath
	if filepath == "" {
		cLogger = log.New(os.Stdout, "", 0)
	} else {
		// Open log file
		logFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
			return err
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

// Function LogInfo writes a ERROR entry in the log
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

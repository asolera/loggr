package loggr

import (
	"strings"
)

var allowedLogTypes map[string]bool
var defaultLogTypes []string
var dateFormat string

func init() {
	allowedLogTypes = make(map[string]bool)
	defaultLogTypes = []string{"info", "error", "warning", "notice", "fatal", "debug", "sql", "timer"}
	dateFormat = "2006-01-02 15:04:05"

	reset(false)
	setDefaults()
}

func reset(value bool) {
	for _, logType := range defaultLogTypes {
		allowedLogTypes[logType] = value
	}
}

func setDefaults() {
	allowedLogTypes["info"] = true
	allowedLogTypes["error"] = true
}

// SetAllowedLogs defines which type of logs should be printed
func SetAllowedLogs(allowedLogTypesString string) {
	if allowedLogTypesString == "all" {
		reset(true)
	} else if allowedLogTypesString == "none" {
		reset(false)
	} else {
		reset(false)
		allowedLogTypesSlice := strings.Split(allowedLogTypesString, "|")

		for _, logType := range allowedLogTypesSlice {
			if _, exists := allowedLogTypes[logType]; exists {
				allowedLogTypes[logType] = true
			}
		}
	}
}

// SetCustomLog allows to configure a custom log type (disabled by default).
// It must be called before "SetAllowedLogs" method
func SetCustomLog(logType string) {
	defaultLogTypes = append(defaultLogTypes, logType)
}

// SetDateFormat defines a custom format for date/time logging (2006-01-02 15:04:05 by default).
func SetDateFormat(layout string) {
	dateFormat = layout
}

// GetLogTypeStatus returns if a log type is active.
func GetLogTypeStatus(logType string) bool {
	return allowedLogTypes[logType]
}

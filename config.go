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
	allowedLogTypes["fatal"] = true
}

// SetAllowedLogs defines which log types will be printed.
// You can use any combination you want. Log types must be joined by pipe (|).
// You can also use `none` to disable logging and `all` to automatically enable all log types.
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

// SetCustomLog allows to create a custom log type.
// A custom type called `example` will be converted to `[EXAMPLE]`.
// Custom types should also be defined in `SetAllowedLogs` in order to be printed.
// This method must be declared BEFORE `SetAllowedLogs` method.
func SetCustomLog(logType string) {
	defaultLogTypes = append(defaultLogTypes, logType)
}

// SetDateFormat overrides default date/time format.
// Layouts must use the reference time `Mon Jan 2 15:04:05 MST 2006` according to the official documentation.
func SetDateFormat(layout string) {
	dateFormat = layout
}

// getLogTypeStatus returns if a log type is active.
func getLogTypeStatus(logType string) bool {
	return allowedLogTypes[logType]
}

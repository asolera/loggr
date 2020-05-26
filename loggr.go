package loggr

import (
	"fmt"
	"strings"
	"time"
)

var timers map[string]time.Time

func init() {
	timers = make(map[string]time.Time)
}

func printLog(logType string, text string) {
	if getLogTypeStatus(logType) == true {
		logTypePrefix := strings.ToUpper(logType)
		logString := fmt.Sprintf("[%s] %s", logTypePrefix, text)
		fmt.Println(time.Now().Format(dateFormat), logString)
	}
}

// Info prints the message with `[INFO]` prefix.
func Info(text string) {
	printLog("info", text)
}

// Error prints the message with `[ERROR]` prefix.
func Error(text string) {
	printLog("error", text)
}

// Warning prints the message with `[WARNING]` prefix.
func Warning(text string) {
	printLog("warning", text)
}

// Notice prints the message with `[NOTICE]` prefix.
func Notice(text string) {
	printLog("notice", text)
}

// Debug prints the message with `[DEBUG]` prefix.
func Debug(text string) {
	printLog("debug", text)
}

// Fatal prints the message with `[FATAL]` prefix.
func Fatal(text string) {
	printLog("fatal", text)
}

// SQL prints the message with `[SQL]` prefix.
func SQL(query string) {
	printLog("sql", query)
}

// Custom prints the message with a custom prefix.
// In order to work, the custom log type must be set in `SetCustomLog` method and defined in `SetAllowedLogs` method.
// First argument is prefix and second argument is the message. Multiple custom log types are also allowed.
func Custom(logType string, text string) {
	printLog(logType, text)
}

func printTimer(text string) {
	printLog("timer", text)
}

// Line prints a line.
func Line() string {
	return "================================================================================"
}

// TimeStart starts a timer for a given label.
func TimeStart(input ...string) {
	if getLogTypeStatus("timer") == true {
		var label string
		if len(input) == 0 {
			label = "default"
		} else {
			label = input[0]
		}
		timers[label] = time.Now()
	}
}

// TimeEnd ends the timer and prints the difference time in seconds.
// In order to be displayed, `timer` (or `all`) must be set in `SetAllowedLogs` method.
// Multiple timers are also allowed.
func TimeEnd(input ...string) {
	if getLogTypeStatus("timer") == true {
		var label string
		if len(input) == 0 {
			label = "default"
		} else {
			label = input[0]
		}
		endTime := time.Now()
		diff := endTime.Sub(timers[label]).String()
		printTimer(fmt.Sprintf("%s: %s", label, diff))
	}
}

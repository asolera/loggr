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

// Info prints an information text
func Info(text string) {
	printLog("info", text)
}

// Error prints an error text
func Error(text string) {
	printLog("error", text)
}

// Warning prints a warning text
func Warning(text string) {
	printLog("warning", text)
}

// Notice prints a notice text
func Notice(text string) {
	printLog("notice", text)
}

// Debug prints a debug text
func Debug(text string) {
	printLog("debug", text)
}

// Fatal prints a fatal text
func Fatal(text string) {
	printLog("fatal", text)
}

// SQL prints a query
func SQL(query string) {
	printLog("sql", query)
}

// Custom prints a custom type of log
func Custom(logType string, text string) {
	printLog(logType, text)
}

func printTimer(text string) {
	printLog("timer", text)
}

// Line prints a line
func Line() string {
	return "================================================================================"
}

// TimeStart starts a timer for a given label (default if none given)
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

// TimeEnd stops the timer for the given label and prints the difference in seconds ("timer" option must be set through logger.SetAllowedLogs())
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

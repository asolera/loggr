# Loggr

Loggr is a very simplistic Go logging library.

Loggr allows you to:
- define any combination of log levels/types;
- override default date/time format;
- create custom log types;
- set timers just like in javascript.

Loggr supports, by default, the following levels: `INFO`, `NOTICE`, `WARNING`, `ERROR`, `FATAL`, `DEBUG`, `SQL`, `TIMER`.

## Basic Usage

Sample code:
```go
package main

import "github.com/asolera/loggr"

func main() {
	logger.Info(logger.Line())
	logger.Info("Hello world!")
	logger.Info(logger.Line())
}
```

Code output:
```
2020-05-24 16:57:39 [INFO] ================================================================================
2020-05-24 16:57:39 [INFO] Hello world!
2020-05-24 16:57:39 [INFO] ================================================================================
```

## Getting Started

Install Go Loggr library with the following command:

```sh
go get github.com/asolera/go-logger/logger
```

Import in your code:

```sh
import "github.com/asolera/go-logger/logger"
```

Configure log level of your application:

```go
func main() {
	logger.SetAllowedLogs("info|warning|error") // you can also use "none" and "all"
}
```

And start logging:

```go
logger.Info("Hello world!")
```

## Configuration

Docs by example (order matters):

```go
func main() {
	// SetDateFormat
	// Overrides default date/time format.
	// Layouts must use the reference time `Mon Jan 2 15:04:05 MST 2006` according to the official documentation.
	// Default: "2006-01-02 15:04:05".
	loggr.SetDateFormat("02/01/2006 15:04:05") 

	// SetCustomLog
	// Allows you to configure a custom type of log.
	// Custom types are also subject to SetAllowedLogs method.
	loggr.SetCustomLog("test")

	// SetAllowedLogs
	// Defines which log types will be printed.
	// Default: "info|error".
	loggr.SetAllowedLogs("info|debug|error|test") // in this case, "test" is a custom type
	loggr.SetAllowedLogs("none") // disable any logging
	loggr.SetAllowedLogs("all") // all log types will be printed (good for dev environment)
}
```

## Logging

Docs by example:

```go
func main() {
	loggr.SetCustomLog("test")
	loggr.SetAllowedLogs("all") // this is needed in order to print all the messages below

	loggr.Info(loggr.Line()) // prints a line inside an info message
	loggr.Info("Info message...")
	loggr.Notice("Notice message...")
	loggr.Warning("Warning message...")
	loggr.Error("Error message...")
	loggr.Fatal("Fatal message...")
	loggr.Debug("Debug message...")
	loggr.SQL("SELECT something FROM example")
	loggr.Custom("test", "A custom log type message...")
	loggr.Info(loggr.Line())
}
```

Output:
```
2020-05-24 18:14:49 [INFO] ================================================================================
2020-05-24 18:14:49 [INFO] Info message...
2020-05-24 18:14:49 [NOTICE] Notice message...
2020-05-24 18:14:49 [WARNING] Warning message...
2020-05-24 18:14:49 [ERROR] Error message...
2020-05-24 18:14:49 [FATAL] Fatal message...
2020-05-24 18:14:49 [DEBUG] Debug message...
2020-05-24 18:14:49 [SQL] SELECT something FROM example
2020-05-24 18:14:49 [TEST] A custom log type message...
2020-05-24 18:14:49 [INFO] ================================================================================
```

## Timer

Just like in [Javascript](https://www.w3schools.com/jsref/met_console_time.asp), you can set a timer using Loggr.

Usage by example:

```go
func main() {
	loggr.SetAllowedLogs("timer") // "all" can be used too

	loggr.TimeStart() // Without a label, timer will be set as "default" label
	time.Sleep(2 * time.Second) // Delay 2 seconds
	loggr.TimeStart("myTimer") // Multiple timers are also allowed
	time.Sleep(3 * time.Second) // Delay 3 seconds
	loggr.TimeEnd("myTimer") // Use "TimeEnd" to stop the given timer
	loggr.TimeEnd()
}
```

Output:
```
2020-05-24 18:13:34 [TIMER] myTimer: 3.0002872s
2020-05-24 18:13:34 [TIMER] default: 5.0571503s
```

## Contributing

1. Fork it
1. Download your fork to your PC (`git clone https://github.com/asolera/loggr && cd loggr`)
1. Create your feature branch (git checkout -b my-new-feature)
1. Make changes and add them (git add .)
1. Commit your changes (git commit -m 'Add some feature')
1. Push to the branch (git push origin my-new-feature)
1. Create new pull request

## Author

Andrew Solera - andrewsolera@gmail.com
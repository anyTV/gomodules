package logger

import (
	"fmt"
	"log"
	"os"
)

// TODO: replace implementation with log/slog
// See https://stackoverflow.com/a/76867161 for possible implementation

// reset
var colorReset = "\033[0m"

// COLORS
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"


// Reserved Color Variables
var colorMagenta = "\033[35m"
var colorCyan = "\033[36m"
var colorGray = "\033[37m"
var colorWhite = "\033[97m"

// levelType
type levelType int8

const (
	DEBUG levelType = iota - 1 // -1
	INFO                       // 0
	WARN                       // 1
	ERROR                      // 2
	FATAL                      // 3
)

func (l levelType) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "ERROR"
	default:
		return "DEBUG"
	}
}

func (l levelType) color() string {
	switch l {
	case DEBUG:
		return colorBlue
	case INFO:
		return colorGreen
	case WARN:
		return colorYellow
	case ERROR:
		return colorRed
	case FATAL:
		return colorRed
	default:
		return colorGreen
	}
}

// module instance
var globalLogger = New("sys", INFO)

func SetLevel(v levelType) {
	globalLogger.SetLevel(v)
}

func SetContext(c string) {
	globalLogger.SetContext(c)
}

// START - Printf

func Debugf(format string, v ...any) {
	globalLogger.Debugf(format, v...)
}

func Infof(format string, v ...any) {
	globalLogger.Infof(format, v...)
}

func Warnf(format string, v ...any) {
	globalLogger.Warnf(format, v...)
}

func Errorf(format string, v ...any) {
	globalLogger.Errorf(format, v...)
}

// Fatalf
//
// Prints log, then calls os.Exit(1)
func Fatalf(format string, v ...any) {
	globalLogger.Fatalf(format, v...)
}

// END - Printf

// START - Println

func Debugln(v ...any) {
	globalLogger.Debugln(v...)
}
func Infoln(v ...any) {
	globalLogger.Infoln(v...)
}
func Warnln(v ...any) {
	globalLogger.Warnln(v...)
}
func Errorln(v ...any) {
	globalLogger.Errorln(v...)
}
func Fatalln(v ...any) {
	globalLogger.Fatalln(v...)
}

// END - Println

// START - Print

func Debug(v ...any) {
	globalLogger.Debug(v...)
}
func Info(v ...any) {
	globalLogger.Info(v...)
}
func Warn(v ...any) {
	globalLogger.Warn(v...)
}
func Error(v ...any) {
	globalLogger.Error(v...)
}
func Fatal(v ...any) {
	globalLogger.Fatal(v...)
}

// END - Print

// START - logStruct

type logStruct struct {
	ctx         string
	logInstance *log.Logger
	level       levelType
}

func (ll logStruct) printf(lvl levelType, format string, v ...any) {
	if lvl < ll.level {
		return
	}

	ll.logInstance.Printf(lvl.color() + "["+ll.ctx+"] "+format+ colorReset, v...)
}


func (ll logStruct) Debugf(format string, v ...any) {
	ll.printf(DEBUG, format, v...)
}

func (ll logStruct) Infof(format string, v ...any) {
	ll.printf(INFO, format, v...)
}

func (ll logStruct) Warnf(format string, v ...any) {
	ll.printf(WARN, format, v...)
}

func (ll logStruct) Errorf(format string, v ...any) {
	ll.printf(ERROR, format, v...)
}

// Fatalf
//
// Prints with Printf, followed by os.Exit(1)
func (ll logStruct) Fatalf(format string, v ...any) {
	ll.printf(FATAL, format, v...)
	os.Exit(1)
}


func (ll logStruct) println(lvl levelType, v ...any) {
	if lvl < ll.level {
		return
	}

	ll.logInstance.Println(fmt.Sprintf(lvl.color() + "[%s] ", ll.ctx), fmt.Sprint(v...), colorReset)
}

func (ll logStruct) Debugln(v ...any) {
	ll.println(DEBUG, v...)
}
func (ll logStruct) Infoln(v ...any) {
	ll.println(INFO, v...)
}
func (ll logStruct) Warnln(v ...any) {
	ll.println(WARN, v...)
}
func (ll logStruct) Errorln(v ...any) {
	ll.println(ERROR, v...)
}

// Fatalln
//
// Prints with Println, followed by os.Exit(1)
func (ll logStruct) Fatalln(v ...any) {
	ll.println(FATAL, v...)
	os.Exit(1)
}


func (ll logStruct) print(lvl levelType, v ...any) {
	if lvl < ll.level {
		return
	}

	ll.logInstance.Print(fmt.Sprintf(lvl.color() + "[%s] ", ll.ctx), fmt.Sprint(v...), colorReset)
}

func (ll logStruct) Debug(v ...any) {
	ll.print(DEBUG, v...)
}
func (ll logStruct) Info(v ...any) {
	ll.print(INFO, v...)
}
func (ll logStruct) Warn(v ...any) {
	ll.print(WARN, v...)
}
func (ll logStruct) Error(v ...any) {
	ll.print(ERROR, v...)
}
// Fatal
// Prints with Print, followed by os.Exit(1)
func (ll logStruct) Fatal(v ...any) {
	ll.print(FATAL, v...)
	os.Exit(1)
}

func (ll logStruct) SetLevel(v levelType) {
	ll.level = v
}

func (ll logStruct) SetContext(c string) {
	ll.ctx = c
}

func New(ctx string, l levelType) logStruct {
	return logStruct{ctx, log.New(os.Stderr, "", log.LstdFlags), l}
}

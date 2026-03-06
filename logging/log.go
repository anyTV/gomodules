package logger

import (
	"fmt"
	"log"
	"os"
)

// TODO: replace implementation with log/slog
// See https://stackoverflow.com/a/76867161 for possible implementation

// module instance
var globalLogger = New("sys", INFO)

func SetLevel(v levelType) {
	globalLogger.SetLevel(v)
}

func GetLevel() levelType {
	return globalLogger.level
}

func SetContext(c string) {
	globalLogger.SetContext(c)
}

// START - Printf

func Verbosef(format string, v ...any) {
	globalLogger.Verbosef(format, v...)
}

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

func Verboseln(v ...any) {
	globalLogger.Verboseln(v...)
}

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

func Verbose(v ...any) {
	globalLogger.Verbose(v...)
}
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

type LogStruct struct {
	ctx         string
	logInstance *log.Logger
	level       levelType
}

func (ll LogStruct) printf(lvl levelType, format string, v ...any) {
	if lvl < ll.level {
		return
	}

	ll.logInstance.Printf(lvl.color()+"["+ll.ctx+"] "+format+colorReset, v...)
}

func (ll LogStruct) Verbosef(format string, v ...any) {
	ll.printf(VERBOSE, format, v...)
}

func (ll LogStruct) Debugf(format string, v ...any) {
	ll.printf(DEBUG, format, v...)
}

func (ll LogStruct) Infof(format string, v ...any) {
	ll.printf(INFO, format, v...)
}

func (ll LogStruct) Warnf(format string, v ...any) {
	ll.printf(WARN, format, v...)
}

func (ll LogStruct) Errorf(format string, v ...any) {
	ll.printf(ERROR, format, v...)
}

// Fatalf
//
// Prints with Printf, followed by os.Exit(1)
func (ll LogStruct) Fatalf(format string, v ...any) {
	ll.printf(FATAL, format, v...)
	os.Exit(1)
}

func (ll LogStruct) println(lvl levelType, v ...any) {
	if lvl < ll.level {
		return
	}

	ll.logInstance.Println(fmt.Sprintf(lvl.color()+"[%s] ", ll.ctx), fmt.Sprint(v...), colorReset)
}

func (ll LogStruct) Verboseln(v ...any) {
	ll.println(VERBOSE, v...)
}

func (ll LogStruct) Debugln(v ...any) {
	ll.println(DEBUG, v...)
}
func (ll LogStruct) Infoln(v ...any) {
	ll.println(INFO, v...)
}
func (ll LogStruct) Warnln(v ...any) {
	ll.println(WARN, v...)
}
func (ll LogStruct) Errorln(v ...any) {
	ll.println(ERROR, v...)
}

// Fatalln
//
// Prints with Println, followed by os.Exit(1)
func (ll LogStruct) Fatalln(v ...any) {
	ll.println(FATAL, v...)
	os.Exit(1)
}

func (ll LogStruct) print(lvl levelType, v ...any) {
	if lvl < ll.level {
		return
	}

	ll.logInstance.Print(fmt.Sprintf(lvl.color()+"[%s] ", ll.ctx), fmt.Sprint(v...), colorReset)
}

func (ll LogStruct) Verbose(v ...any) {
	ll.print(VERBOSE, v...)
}

func (ll LogStruct) Debug(v ...any) {
	ll.print(DEBUG, v...)
}
func (ll LogStruct) Info(v ...any) {
	ll.print(INFO, v...)
}
func (ll LogStruct) Warn(v ...any) {
	ll.print(WARN, v...)
}
func (ll LogStruct) Error(v ...any) {
	ll.print(ERROR, v...)
}

// Fatal
// Prints with Print, followed by os.Exit(1)
func (ll LogStruct) Fatal(v ...any) {
	ll.print(FATAL, v...)
	os.Exit(1)
}

func (ll *LogStruct) SetLevel(v levelType) {
	ll.level = v
}

func (ll *LogStruct) SetContext(c string) {
	ll.ctx = c
}

func (ll LogStruct) GetLevel() levelType {
	return ll.level
}

func (ll LogStruct) GetContext() string {
	return ll.ctx
}

func New(ctx string, l levelType) LogStruct {
	return LogStruct{ctx, log.New(os.Stderr, "", log.LstdFlags), l}
}

package logger

import (
	"fmt"
	"log"
	"os"
)

// TODO: replace implementation with log/slog
// See https://stackoverflow.com/a/76867161 for possible implementation

// reset
var rst = "\033[0m"

// colors
var red = "\033[31m"
var grn = "\033[32m"
var ylw = "\033[33m"
var blu = "\033[34m"
var mag = "\033[35m"
var cya = "\033[36m"
var gry = "\033[37m"
var wht = "\033[97m"

var l = *log.New(os.Stderr, "", log.LstdFlags)

func Debugf(format string, v ...any) { l.Printf(blu+format+rst, v...) }
func Infof(format string, v ...any)  { l.Printf(grn+format+rst, v...) }
func Warnf(format string, v ...any)  { l.Printf(ylw+format+rst, v...) }
func Errorf(format string, v ...any) { l.Printf(red+format+rst, v...) }
func Fatalf(format string, v ...any) { l.Fatalf(red+format+rst, v...) }

func Debugln(v ...any) { l.Print(blu, fmt.Sprint(v...), rst) }
func Infoln(v ...any)  { l.Print(grn, fmt.Sprint(v...), rst) }
func Warnln(v ...any)  { l.Print(ylw, fmt.Sprint(v...), rst) }
func Errorln(v ...any) { l.Print(red, fmt.Sprint(v...), rst) }
func Fatalln(v ...any) { l.Fatal(red, fmt.Sprint(v...), rst) }

func Debug(v ...any) { l.Print(blu, fmt.Sprint(v...), rst) }
func Info(v ...any)  { l.Print(grn, fmt.Sprint(v...), rst) }
func Warn(v ...any)  { l.Print(ylw, fmt.Sprint(v...), rst) }
func Error(v ...any) { l.Print(red, fmt.Sprint(v...), rst) }
func Fatal(v ...any) { l.Fatal(red, fmt.Sprint(v...), rst) }

type L struct{ ctx string }

func (ll L) Debugf(format string, v ...any) { Debugf("["+ll.ctx+"]"+" "+format, v...) }
func (ll L) Infof(format string, v ...any)  { Infof("["+ll.ctx+"]"+" "+format, v...) }
func (ll L) Warnf(format string, v ...any)  { Warnf("["+ll.ctx+"]"+" "+format, v...) }
func (ll L) Errorf(format string, v ...any) { Errorf("["+ll.ctx+"]"+" "+format, v...) }
func (ll L) Fatalf(format string, v ...any) { Fatalf("["+ll.ctx+"]"+" "+format, v...) }

func (ll L) Debugln(v ...any) { Debugln("[", ll.ctx, "]", " ", fmt.Sprint(v...)) }
func (ll L) Infoln(v ...any)  { Infoln("[", ll.ctx, "]", " ", fmt.Sprint(v...)) }
func (ll L) Warnln(v ...any)  { Warnln("[", ll.ctx, "]", " ", fmt.Sprint(v...)) }
func (ll L) Errorln(v ...any) { Errorln("[", ll.ctx, "]", " ", fmt.Sprint(v...)) }
func (ll L) Fatalln(v ...any) { Fatalln("[", ll.ctx, "]", " ", fmt.Sprint(v...)) }

func (ll L) Debug(v ...any) { Debug("[", ll.ctx, "]", " ", fmt.Sprint(v...)) }
func (ll L) Info(v ...any)  { Info("[", ll.ctx, "]", " ", fmt.Sprint(v...)) }
func (ll L) Warn(v ...any)  { Warn("[", ll.ctx, "]", " ", fmt.Sprint(v...)) }
func (ll L) Error(v ...any) { Error("[", ll.ctx, "]", " ", fmt.Sprint(v...)) }
func (ll L) Fatal(v ...any) { Fatal("[", ll.ctx, "]", " ", fmt.Sprint(v...)) }

func New(ctx string) L {
	return L{ctx}
}

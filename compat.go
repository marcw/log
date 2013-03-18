// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"fmt"
	"os"
)

var stdFormatter = &LineFormatter{LineFormat: "%datetime% %level_name%: %message%"}
var stdHandler = &writeCloserHandler{wc: os.Stderr, Handler: &Handler{Level: DEBUG, Formatter: stdFormatter}}
var std = &Logger{Name: "", handlers: []HandlerInterface{stdHandler}, processors: []Processor{}}

// Fatal is equivalent to a call to Print followed by a call to os.Exit(1)
func Fatal(v ...interface{}) {
	Print(v)
	os.Exit(1)
}

// Fatal is equivalent to a call to Printf followed by a call to os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	Printf(format, v)
	os.Exit(1)
}

// Fatalln is equivalent to a call to Println followed by a call to os.Exit(1)
func Fatalln(v ...interface{}) {
	Println(v)
	os.Exit(1)
}

// Panic is equivalent to a call to Print followed by a call to panic
func Panic(v ...interface{}) {
	s := fmt.Sprint(v)
	Print(s)
	panic(s)
}

// Panicf is equivalent to a call to Printf followed by a call to panic
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	Print(s)
	panic(s)
}

// Panicln is equivalent to a call to Println followed by a call to panic
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	Print(s)
	panic(s)
}

// Print calls Debug in an instance of Logger where the only handler outputs to Stderr
func Print(v ...interface{}) {
	std.Debug(v...)
}

// Printf calls Debug in an instance of Logger where the only handler outputs to Stderr
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	std.Debug(fmt.Sprintf(format, v...))
}

// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	std.Debug(fmt.Sprintln(v...))
}

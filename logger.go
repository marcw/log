// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// A `log/syslog` interface to `log` with a user MultiLogger implementation
package gogol

import (
	"log"
)

type Logger interface {
	Alert(m string) (err error)
	Crit(m string) (err error)
	Debug(m string) (err error)
	Emerg(m string) (err error)
	Err(m string) (err error)
	Info(m string) (err error)
	Notice(m string) (err error)
	Warning(m string) (err error)
	Write(m string) (err error)
}

// StdLogger is a proxy wrapping a standard log.Logger instance
type StdLogger struct {
	logger *log.Logger
}

func NewStdLogger(l *log.Logger) *StdLogger {
	return &StdLogger{logger: l}
}

func (l *StdLogger) Write(m string) (err error) {
	l.logger.Println(m)
	return nil
}

// Prefix m with WARNING and write it
func (l *StdLogger) Warning(m string) (err error) {
	l.logger.Println("WARNING", m)
	return nil
}

// Prefix m with NOTICE and write it
func (l *StdLogger) Notice(m string) (err error) {
	l.logger.Println("NOTICE", m)
	return nil
}

// Prefix m with INFO and write it
func (l *StdLogger) Info(m string) (err error) {
	l.logger.Println("INFO", m)
	return nil
}

// Prefix m with ERR and write it
func (l *StdLogger) Err(m string) (err error) {
	l.logger.Println("ERR", m)
	return nil
}

// Prefix m with EMERG and write it
func (l *StdLogger) Emerg(m string) (err error) {
	l.logger.Println("EMERG", m)
	return nil
}

// Prefix m with DEBUG and write it
func (l *StdLogger) Debug(m string) (err error) {
	l.logger.Println("DEBUG", m)
	return nil
}

// Prefix m with CRIT and write it
func (l *StdLogger) Crit(m string) (err error) {
	l.logger.Println("CRIT", m)
	return nil
}

// Prefix m with ALERT and write it
func (l *StdLogger) Alert(m string) (err error) {
	l.logger.Println("ALERT", m)
	return nil
}

// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"fmt"
)

// A logger will log records transformed by the default processors to a collection of handlers
type Logger struct {
	Name       string
	handlers   []Handler
	processors []Processor
}

// Instanciates a new logger with specified name, handlers and processors
func NewLogger(name string) *Logger {
	return &Logger{Name: name, handlers: []Handler{}, processors: []Processor{}}
}

// Push a handler to the handlers stack
func (l *Logger) PushHandler(h Handler) {
	handlers := make([]Handler, len(l.handlers))
	copy(handlers, l.handlers)

	l.handlers = []Handler{h}
	l.handlers = append(l.handlers, handlers...)
}

// Pop a handler from the handlers stack
func (l *Logger) PopHandler() {
	if len(l.handlers) > 0 {
		l.handlers = l.handlers[1:len(l.handlers)]
		return
	}

	panic("Handlers stack is empty")
}

// Push a processor to the processor stack
func (l *Logger) PushProcessor(p Processor) {
	processors := make([]Processor, len(l.processors))
	copy(processors, l.processors)

	l.processors = []Processor{p}
	l.processors = append(l.processors, processors...)
}

// Pop a processor from the processor stack
func (l *Logger) PopProcessor() {
	if len(l.processors) > 0 {
		l.processors = l.processors[1:len(l.processors)]
		return
	}

	panic("Processors stack is empty")
}

// Log string with specified severity
func (l *Logger) AddRecord(level Severity, message string) {
	r := newRecord(level, l.Name, message)

	if !l.IsHandling(level) {
		return
	}

	for k := range l.processors {
		l.processors[k].Process(r)
	}

	for k := range l.handlers {
		if l.handlers[k].IsHandling(level) {
			l.handlers[k].Handle(r)
		}
	}
}

// Log string with the DEBUG level
func (l *Logger) AddDebug(message string) {
	l.AddRecord(DEBUG, message)
}

// Log string with the INFO level
func (l *Logger) AddInfo(message string) {
	l.AddRecord(INFO, message)
}

// Log string with the NOTICE level
func (l *Logger) AddNotice(message string) {
	l.AddRecord(NOTICE, message)
}

// Log string with the WARNING level
func (l *Logger) AddWarning(message string) {
	l.AddRecord(WARNING, message)
}

// Log string with the ERROR level
func (l *Logger) AddError(message string) {
	l.AddRecord(ERROR, message)
}

// Log string with the CRITICAL level
func (l *Logger) AddCritical(message string) {
	l.AddRecord(CRITICAL, message)
}

// Log string with the ALERT level
func (l *Logger) AddAlert(message string) {
	l.AddRecord(ALERT, message)
}

// Log string with the EMERGENCY level
func (l *Logger) AddEmergency(message string) {
	l.AddRecord(EMERGENCY, message)
}

// Log parameters with the DEBUG level
func (l *Logger) Debug(v ...interface{}) {
	l.AddDebug(fmt.Sprint(v...))
}

// Log parameters with the INFO level
func (l *Logger) Info(v ...interface{}) {
	l.AddInfo(fmt.Sprint(v...))
}

// Log parameters with the NOTICE level
func (l *Logger) Notice(v ...interface{}) {
	l.AddNotice(fmt.Sprint(v...))
}

// Log parameters with the WARNING level
func (l *Logger) Warning(v ...interface{}) {
	l.AddWarning(fmt.Sprint(v...))
}

// Log parameters with the ERROR level
func (l *Logger) Error(v ...interface{}) {
	l.AddError(fmt.Sprint(v...))
}

// Log parameters with the CRITICAL level
func (l *Logger) Critical(v ...interface{}) {
	l.AddCritical(fmt.Sprint(v...))
}

// Log parameters with the ALERT level
func (l *Logger) Alert(v ...interface{}) {
	l.AddAlert(fmt.Sprint(v...))
}

// Log parameters with the EMERGENCY level
func (l *Logger) Emergency(v ...interface{}) {
	l.AddEmergency(fmt.Sprint(v...))
}

// Returns true if a Handler can handle this severity level
func (l *Logger) IsHandling(level Severity) bool {
	for k := range l.handlers {
		if l.handlers[k].IsHandling(level) {
			return true
		}
	}

	return false
}

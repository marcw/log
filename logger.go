// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"fmt"
)

type Logger struct {
	Name       string
	handlers   []Handler
	processors []Processor
}

func NewLogger(name string, handlers []Handler, processors []Processor) *Logger {
	return &Logger{Name: name, handlers: handlers, processors: processors}
}

func (l *Logger) PushHandler(h Handler) {
	handlers := make([]Handler, len(l.handlers))
	copy(handlers, l.handlers)

	l.handlers = []Handler{h}
	l.handlers = append(l.handlers, handlers...)
}

func (l *Logger) PopHandler() {
	if len(l.handlers) > 0 {
		l.handlers = l.handlers[1:len(l.handlers)]
        return
	}

	panic("Handlers stack is empty")
}

func (l *Logger) PushProcessor(p Processor) {
	processors := make([]Processor, len(l.processors))
	copy(processors, l.processors)

	l.processors = []Processor{p}
	l.processors = append(l.processors, processors...)
}

func (l *Logger) PopProcessor() {
	if len(l.processors) > 0 {
		l.processors = l.processors[1:len(l.processors)]
        return
	}

	panic("Processors stack is empty")
}

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

func (l *Logger) AddDebug(message string) {
	l.AddRecord(DEBUG, message)
}

func (l *Logger) AddInfo(message string) {
	l.AddRecord(INFO, message)
}

func (l *Logger) AddNotice(message string) {
	l.AddRecord(NOTICE, message)
}

func (l *Logger) AddWarning(message string) {
	l.AddRecord(WARNING, message)
}

func (l *Logger) AddError(message string) {
	l.AddRecord(ERROR, message)
}

func (l *Logger) AddCritical(message string) {
	l.AddRecord(CRITICAL, message)
}

func (l *Logger) AddAlert(message string) {
	l.AddRecord(ALERT, message)
}

func (l *Logger) AddEmergency(message string) {
	l.AddRecord(EMERGENCY, message)
}

func (l *Logger) Debug(v ...interface{}) {
	l.AddDebug(fmt.Sprint(v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.AddInfo(fmt.Sprint(v...))
}

func (l *Logger) Notice(v ...interface{}) {
	l.AddNotice(fmt.Sprint(v...))
}

func (l *Logger) Warning(v ...interface{}) {
	l.AddWarning(fmt.Sprint(v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.AddError(fmt.Sprint(v...))
}

func (l *Logger) Critical(v ...interface{}) {
	l.AddCritical(fmt.Sprint(v...))
}

func (l *Logger) Alert(v ...interface{}) {
	l.AddAlert(fmt.Sprint(v...))
}

func (l *Logger) Emergency(v ...interface{}) {
	l.AddEmergency(fmt.Sprint(v...))
}

func (l *Logger) IsHandling(level Severity) bool {
	for k := range l.handlers {
		if l.handlers[k].IsHandling(level) {
			return true
		}
	}

	return false
}

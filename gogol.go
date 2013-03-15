// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"fmt"
	"time"
)

type Severity int

const (
	EMERGENCY Severity = iota // Emergency: system is unusable
	ALERT                     //  Alert: action must be taken immediately
	CRITICAL                  // Critical: critical conditions
	ERROR                     // Error: error conditions
	WARNING                   //  Warning: warning conditions
	NOTICE                    // Notice: normal but significant condition
	INFO                      // Informational: informational messages
	DEBUG                     //Debug: debug-level messages
)

var Severities = map[Severity]string{
	DEBUG:     "DEBUG",
	INFO:      "INFO",
	NOTICE:    "NOTICE",
	WARNING:   "WARNING",
	ERROR:     "ERROR",
	CRITICAL:  "CRITICAL",
	ALERT:     "ALERT",
	EMERGENCY: "EMERGENCY"}

type Record struct {
	Message   string
	Formatted string
	Level     Severity
	LevelName string
	Channel   string
	Time      string
	Extra     Extra
}

type Extra map[string]interface{}

func (e Extra) Normalize() map[string]string {
	normalized := make(map[string]string)

	for k, v := range e {
		switch x := v.(type) {
		case time.Time:
			normalized[k] = x.Format(time.RFC3339Nano)
			break
		case fmt.Stringer:
			normalized[k] = x.String()
			break
		default:
			normalized[k] = fmt.Sprint(v)
		}
	}

	return normalized
}

func (e Extra) String() string {
	// TODO: This is a really shitty way to convert this map to a string
	return fmt.Sprint(e.Normalize())
}

func newRecord(level Severity, channel, message string) *Record {
	return &Record{
		Message:   message,
		Level:     level,
		LevelName: Severities[level],
		Channel:   channel,
		Time:      time.Now().Format(time.RFC3339Nano),
		Extra:     make(Extra)}
}

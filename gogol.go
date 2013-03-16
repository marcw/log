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
	// From /usr/include/sys/syslog.h. and RFC5424
	EMERGENCY Severity = iota // Emergency: system is unusable
	ALERT                     // Alert: action must be taken immediately
	CRITICAL                  // Critical: critical conditions
	ERROR                     // Error: error conditions
	WARNING                   // Warning: warning conditions
	NOTICE                    // Notice: normal but significant condition
	INFO                      // Informational: informational messages
	DEBUG                     // Debug: debug-level messages
)

// Textual translation of severities
var Severities = map[Severity]string{
	DEBUG:     "DEBUG",
	INFO:      "INFO",
	NOTICE:    "NOTICE",
	WARNING:   "WARNING",
	ERROR:     "ERROR",
	CRITICAL:  "CRITICAL",
	ALERT:     "ALERT",
	EMERGENCY: "EMERGENCY"}

// A record is a log message at a given time
type Record struct {
	Message   string   // Text message of the log
	Formatted string   // Formatted version of the log (once all processors and formatters have done their jobs)
	Level     Severity // Severity level
	LevelName string   // Severity name
	Channel   string   // Logger's name
	Time      string   // Creation date formated to time.RFC3339Nano
	Extra     Extra    // Extra values that can be added by Processors
}

type Extra map[string]interface{}

// Normalize data to a map of string
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

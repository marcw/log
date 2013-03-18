// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"strings"
)

// A formatter formats the record before being sent to a Handler
// The result of the formatting MUST be in Record.Formatted
type Formatter interface {
	Format(*Record)
}

const (
	LineFormatSimple  string = "[%datetime%] %channel%.%level_name%: %message% %extra%\n"
	LineFormatMinimal string = "%channel%.%level_name%: %message%\n"
)

// A line formatter formats a Record into a line of text.
// Available formats are LINE_FORMAT_SIMPLE, LINE_FORMAT_MINIMAL or you can make your own
// with these fields:
// %datetime%: Record's creation date in the time.RFC3339Nano format
// $channel%: logger.Name
// %level_name%: Severity's name (DEBUG, WARNING, ...)
// %message%: Message text
// %extra%: All extra values, generally added by Processors
type LineFormatter struct {
	LineFormat string
}

// Instantiates a new LineFormatter with the LINE_FORMAT_MINIMAL format
func NewMinimalLineFormatter() Formatter {
	return &LineFormatter{LineFormat: LineFormatMinimal}
}

// Instantiates a new LineFormatter with the LINE_FORMAT_SIMPLE format
func NewSimpleLineFormatter() Formatter {
	return &LineFormatter{LineFormat: LineFormatSimple}
}

// Format the Record r with f.LineFormat
func (f *LineFormatter) Format(r *Record) {
	replacer := strings.NewReplacer(
		"%datetime%", r.Time,
		"%channel%", r.Channel,
		"%level_name%", r.LevelName,
		"%message%", r.Message,
		"%extra%", r.Extra.String())

	r.Formatted = replacer.Replace(f.LineFormat)
}

// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"strings"
)

type Formatter interface {
	Format(*Record)
}

const LINE_FORMAT_SIMPLE = "[%datetime%] %channel%.%level_name%: %message% %extra%\n"

type LineFormatter struct {
	LineFormat string
}

func NewSimpleLineFormatter() Formatter {
	return &LineFormatter{LineFormat: LINE_FORMAT_SIMPLE}
}

func (f *LineFormatter) Format(r *Record) {
	replacer := strings.NewReplacer(
		"%datetime%", r.Time,
		"%channel%", r.Channel,
		"%level_name%", r.LevelName,
		"%message%", r.Message,
		"%extra%", r.Extra.String())

	r.Formatted = replacer.Replace(f.LineFormat)
}

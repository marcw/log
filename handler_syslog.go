// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"log/syslog"
)

type syslogHandler struct {
	w *syslog.Writer
	*Handler
}

// Instantiates a new Handler which will write on syslog when level is reached.
func NewSyslogHandler(sw *syslog.Writer, level Severity) HandlerInterface {
	return &syslogHandler{sw, &Handler{Level: level}}
}

func (sh *syslogHandler) Close() {
	sh.w.Close()
}

func (sh *syslogHandler) Handle(r *Record) {
	sh.Handler.Prepare(r)
	switch r.Level {
	case DEBUG:
		sh.w.Debug(r.Formatted)
		break
	case INFO:
		sh.w.Info(r.Formatted)
		break
	case NOTICE:
		sh.w.Notice(r.Formatted)
		break
	case WARNING:
		sh.w.Warning(r.Formatted)
		break
	case ERROR:
		sh.w.Err(r.Formatted)
		break
	case CRITICAL:
		sh.w.Crit(r.Formatted)
		break
	case ALERT:
		sh.w.Alert(r.Formatted)
		break
	case EMERGENCY:
		sh.w.Emerg(r.Formatted)
		break
	default:
		panic("Unreachable")
	}
}

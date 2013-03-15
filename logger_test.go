// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"bytes"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	buffer := new(bytes.Buffer)
	handler := NewBufferHandler(buffer, NOTICE)
	handler.SetFormatter(&LineFormatter{"%channel%.%level_name%: %message%"})
	logger := NewLogger("test", []Handler{handler}, []Processor{})

	logger.Debug("foobar")
	if "" != buffer.String() {
		t.Error(buffer.String())
	}
	buffer.Reset()
	logger.Info("foobar")
	if "" != buffer.String() {
		t.Error(buffer.String())
	}
	buffer.Reset()
	logger.Notice("foobar")
	if "test.NOTICE: foobar" != buffer.String() {
		t.Error(buffer.String())
	}
	buffer.Reset()
	logger.Warning("foobar")
	if "test.WARNING: foobar" != buffer.String() {
		t.Error(buffer.String())
	}
	buffer.Reset()
	logger.Error("foobar")
	if "test.ERROR: foobar" != buffer.String() {
		t.Error(buffer.String())
	}
	buffer.Reset()
	logger.Critical("foobar")
	if "test.CRITICAL: foobar" != buffer.String() {
		t.Error(buffer.String())
	}
	buffer.Reset()
	logger.Alert("foobar")
	if "test.ALERT: foobar" != buffer.String() {
		t.Error(buffer.String())
	}
	buffer.Reset()
	logger.Emergency("foobar")
	if "test.EMERGENCY: foobar" != buffer.String() {
		t.Error(buffer.String())
	}
}

func TestLoggerPushPopProcessors(t *testing.T) {
	p1 := NewProcessor(func(r *Record) {}) // Will be called last
	p2 := NewProcessor(func(r *Record) {})

	l := NewLogger("test", []Handler{}, []Processor{})
	l.PushProcessor(p1)
	l.PushProcessor(p2)

	if len(l.processors) != 2 {
		t.Error()
	}
	if l.processors[0] != p2 {
		t.Error()
	}
	if l.processors[1] != p1 {
		t.Error()
	}
	l.PopProcessor()
	if len(l.processors) != 1 {
		t.Error()
	}
	if l.processors[0] != p1 {
		t.Error()
	}
}

func TestLoggerPushPopHandler(t *testing.T) {
	buffer := new(bytes.Buffer)
	h1 := NewBufferHandler(buffer, DEBUG)
	h2 := NewBufferHandler(buffer, DEBUG)
	l := &Logger{}
	l.PushHandler(h1)
	l.PushHandler(h2)

	if len(l.handlers) != 2 {
		t.Error()
	}
	if l.handlers[0] != h2 {
		t.Error()
	}
	if l.handlers[1] != h1 {
		t.Error()
	}
	l.PopHandler()
	if len(l.handlers) != 1 {
		t.Error()
	}
	if l.handlers[0] != h1 {
		t.Error()
	}
}

func TestLoggerWithDefaultProcessor(t *testing.T) {
	buffer := new(bytes.Buffer)
	h1 := NewBufferHandler(buffer, DEBUG)
	h1.SetFormatter(&LineFormatter{"%message% %extra%"})
	l := NewLogger("channel", []Handler{}, []Processor{})
	l.PushHandler(h1)
	p := NewProcessor(func(r *Record) {
		r.Extra["go.date"] = time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)
	})
	l.PushProcessor(p)
	l.Debug("foobar")
	if "foobar map[go.date:2009-11-10T23:00:00Z]" != buffer.String() {
		t.Error()
	}
}

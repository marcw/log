// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"testing"
)

func TestIsHandling(t *testing.T) {
	h := &handler{level: EMERGENCY}
	if h.IsHandling(DEBUG) {
		t.Error()
	}
	if h.IsHandling(INFO) {
		t.Error()
	}
	if h.IsHandling(NOTICE) {
		t.Error()
	}
	if h.IsHandling(WARNING) {
		t.Error()
	}
	if h.IsHandling(ERROR) {
		t.Error()
	}
	if h.IsHandling(CRITICAL) {
		t.Error()
	}
	if h.IsHandling(ALERT) {
		t.Error()
	}
	if !h.IsHandling(EMERGENCY) {
		t.Error()
	}

	h = &handler{level: DEBUG}
	if !h.IsHandling(DEBUG) {
		t.Error()
	}
	if !h.IsHandling(INFO) {
		t.Error()
	}
	if !h.IsHandling(NOTICE) {
		t.Error()
	}
	if !h.IsHandling(WARNING) {
		t.Error()
	}
	if !h.IsHandling(ERROR) {
		t.Error()
	}
	if !h.IsHandling(CRITICAL) {
		t.Error()
	}
	if !h.IsHandling(ALERT) {
		t.Error()
	}
	if !h.IsHandling(EMERGENCY) {
		t.Error()
	}
}

func TestPushAndPopProcessor(t *testing.T) {
	h := &handler{}
	p1 := NewProcessor(func(r *Record) {})
	p2 := NewProcessor(func(r *Record) {})
	h.PushProcessor(p1)
	h.PushProcessor(p2)

	if len(h.processors) != 2 {
		t.Error()
	}
	if h.processors[0] != p2 {
		t.Error()
	}
	if h.processors[1] != p1 {
		t.Error()
	}
	h.PopProcessor()
	if len(h.processors) != 1 {
		t.Error()
	}
	if h.processors[0] != p1 {
		t.Error()
	}
}

func TestPrepare(t *testing.T) {
	h := &handler{level: DEBUG}
	p1 := NewProcessor(func(r *Record) { r.Message = "p1" }) // Will be called last
	p2 := NewProcessor(func(r *Record) { r.Message = "p2" })
	h.PushProcessor(p1)
	h.PushProcessor(p2)
	h.SetFormatter(&LineFormatter{LineFormat: "%message%"})

	r := &Record{Message: "original"}
	h.Prepare(r)
	if r.Formatted != "p1" {
		t.Error(r.Formatted)
	}
}

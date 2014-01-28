// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package log

import (
	"testing"
)

func TestIsHandling(t *testing.T) {
	h := &Handler{Level: EMERGENCY}
	if h.S(DEBUG) {
		t.Error("handler should not handle DEBUG priority")
	}
	if h.S(INFO) {
		t.Error("handler should not handle INFO priority")
	}
	if h.S(NOTICE) {
		t.Error("handler should not handle EMERGENCY priority")
	}
	if h.S(WARNING) {
		t.Error("handler should not handle WARNING priority")
	}
	if h.S(ERROR) {
		t.Error("handler should not handle ERROR priority")
	}
	if h.S(CRITICAL) {
		t.Error("handler should not handle CRITICAL priority")
	}
	if h.S(ALERT) {
		t.Error("handler should not handle ALERT priority")
	}
	if !h.S(EMERGENCY) {
		t.Error("handler should handle EMERGENCY priority")
	}

	h = &Handler{Level: DEBUG}
	if !h.S(DEBUG) {
		t.Error("handler should handle DEBUG priority")
	}
	if !h.S(INFO) {
		t.Error("handler should handle INFO priority")
	}
	if !h.S(NOTICE) {
		t.Error("handler should handle NOTICE priority")
	}
	if !h.S(WARNING) {
		t.Error("handler should handle WARNING priority")
	}
	if !h.S(ERROR) {
		t.Error("handler should handle ERROR priority")
	}
	if !h.S(CRITICAL) {
		t.Error("handler should handle CRITICAL priority")
	}
	if !h.S(ALERT) {
		t.Error("handler should handle ALERTY priority")
	}
	if !h.S(EMERGENCY) {
		t.Error("handler should handle EMERGENCY priority")
	}
}

func TestPushAndPopProcessor(t *testing.T) {
	h := &Handler{}
	p1 := NewProcessor(func(r *Record) {})
	p2 := NewProcessor(func(r *Record) {})
	h.PushProcessor(p1)
	h.PushProcessor(p2)

	if len(h.Processors) != 2 {
		t.Error("number of processor should be equal to 2")
	}
	if h.Processors[0] != p2 {
		t.Error("processor mismatch")
	}
	if h.Processors[1] != p1 {
		t.Error("processor mismatch")
	}
	h.PopProcessor()
	if len(h.Processors) != 1 {
		t.Error("number of processor should be equal to 1")
	}
	if h.Processors[0] != p1 {
		t.Error("processor mistmatch")
	}
}

func TestPrepare(t *testing.T) {
	h := &Handler{Level: DEBUG}
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

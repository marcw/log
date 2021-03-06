// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package log

import (
	"testing"
	"time"
)

func TestMinimalLineFormatter(t *testing.T) {
	r := newRecord(DEBUG, "foobar", "msg", map[string]string{"foo": "bar"})
	formatter := NewMinimalLineFormatter()
	formatter.Format(r)
	if r.Formatted != "foobar.DEBUG: msg\n" {
		t.Error(r.Formatted)
	}
}

func TestSimpleLineFormatter(t *testing.T) {
	r := newRecord(DEBUG, "foobar", "msg\n", map[string]string{"foo": "bar"})
	r.Time = time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)
	formatter := NewSimpleLineFormatter()
	formatter.Format(r)
	if r.Formatted != "[2009-11-10T23:00:00Z] foobar.DEBUG: msg foo=\"bar\" \n" {
		t.Error(r.Formatted)
	}
}

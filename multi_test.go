// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"bytes"
	"log"
	"log/syslog"
	"testing"
)

func TestMultiLogger(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	sl := NewStdLogger(log.New(b, "", 0))
	l := MultiLogger(sl, sl)

	l.Write([]byte("foobar"))
	if b.String() != "foobar\nfoobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "foobar\n", b.String())
	}

	b.Reset()
	l.Warning("foobar")
	if b.String() != "WARNING foobar\nWARNING foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "WARNING foobar\nWARNING foobar\n", b.String())
	}

	b.Reset()
	l.Notice("foobar")
	if b.String() != "NOTICE foobar\nNOTICE foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "NOTICE foobar\nNOTICE foobar\n", b.String())
	}

	b.Reset()
	l.Info("foobar")
	if b.String() != "INFO foobar\nINFO foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "INFO foobar\nINFO foobar\n", b.String())
	}

	b.Reset()
	l.Err("foobar")
	if b.String() != "ERR foobar\nERR foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "ERR foobar\nERR foobar\n", b.String())
	}

	b.Reset()
	l.Emerg("foobar")
	if b.String() != "EMERG foobar\nEMERG foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "EMERG foobar\nEMERG foobar\n", b.String())
	}

	b.Reset()
	l.Debug("foobar")
	if b.String() != "DEBUG foobar\nDEBUG foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "DEBUG foobar\n", b.String())
	}

	b.Reset()
	l.Crit("foobar")
	if b.String() != "CRIT foobar\nCRIT foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "CRIT foobar\nCRIT foobar\n", b.String())
	}

	b.Reset()
	l.Alert("foobar")
	if b.String() != "ALERT foobar\nALERT foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "ALERT foobar\n", b.String())
	}
}

func TestMultiLoggerAcceptsSyslog(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	sl := NewStdLogger(log.New(b, "", 0))
	s, _ := syslog.New(syslog.LOG_DEBUG, "")
	MultiLogger(sl, s)
}

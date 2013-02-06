// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"bytes"
	"log"
	"testing"
)

func TestStdLogger(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	l := NewStdLogger(log.New(b, "", 0))

	l.Write([]byte("foobar"))
	if b.String() != "foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "foobar\n", b.String())
	}

	b.Reset()
	l.Warning("foobar")
	if b.String() != "WARNING foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "WARNING foobar\n", b.String())
	}

	b.Reset()
	l.Notice("foobar")
	if b.String() != "NOTICE foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "NOTICE foobar\n", b.String())
	}

	b.Reset()
	l.Info("foobar")
	if b.String() != "INFO foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "INFO foobar\n", b.String())
	}

	b.Reset()
	l.Err("foobar")
	if b.String() != "ERR foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "ERR foobar\n", b.String())
	}

	b.Reset()
	l.Emerg("foobar")
	if b.String() != "EMERG foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "EMERG foobar\n", b.String())
	}

	b.Reset()
	l.Debug("foobar")
	if b.String() != "DEBUG foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "DEBUG foobar\n", b.String())
	}

	b.Reset()
	l.Crit("foobar")
	if b.String() != "CRIT foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "CRIT foobar\n", b.String())
	}

	b.Reset()
	l.Alert("foobar")
	if b.String() != "ALERT foobar\n" {
		t.Fatalf("Expected %#v. Found %#v instead.", "ALERT foobar\n", b.String())
	}
}

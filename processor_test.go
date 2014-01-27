// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"runtime"
	"testing"
)

func TestRuntimeProcessor(t *testing.T) {
	r := newRecord(DEBUG, "test", "foobar", map[string]string{"foo": "bar"})
	RuntimeProcessor.Process(r)
	if r.Extra["go.num_cpu"] != runtime.NumCPU() {
		t.Error("go.num_cpu is not correct")
	}
	if r.Extra["go.version"] != runtime.Version() {
		t.Error("go.version is not correct")
	}
	if r.Extra["go.num_goroutines"] != runtime.NumGoroutine() {
		t.Error("go.num_goroutines is not correct")
	}
}

// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"bytes"
	"testing"
)

func TestBufferHandler(t *testing.T) {
	buffer := new(bytes.Buffer)
	h := NewBufferHandler(buffer, WARNING)

	if h.S(DEBUG) {
		t.Error("handler should not handle DEBUG priority")
	}
}

// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"bytes"
)

type bufferHandler struct {
	buffer *bytes.Buffer
	*Handler
}

// Instantiates a new handler that will keep logs in memory
func NewBufferHandler(buffer *bytes.Buffer, level Severity) HandlerInterface {
	return &bufferHandler{buffer, &Handler{Level: level}}
}

func (bh *bufferHandler) Handle(r Record) {
	bh.Handler.Prepare(&r)
	bh.buffer.WriteString(r.Formatted)
}

func (bh *bufferHandler) Close() {
}

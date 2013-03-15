// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"bytes"
)

type bufferHandler struct {
	buffer *bytes.Buffer
	*handler
}

func NewBufferHandler(buffer *bytes.Buffer, level Severity) Handler {
	return &bufferHandler{buffer, &handler{level: level}}
}

func (bh *bufferHandler) Handle(r *Record) {
	bh.handler.Prepare(r)
	bh.buffer.WriteString(r.Formatted)
}

func (bh *bufferHandler) Close() {
}

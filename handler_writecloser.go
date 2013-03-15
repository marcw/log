// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"io"
	"os"
)

type writeCloserHandler struct {
	wc io.WriteCloser
	*handler
}

// Instantiates a new Handler which will write on wc when level is reached.
func NewWriteCloserHandler(wc io.WriteCloser, level Severity) Handler {
	return &writeCloserHandler{wc, &handler{level: level}}
}

// Instantiates a new Handler which will write on Stdout when level is reached.
func NewStdoutHandler(level Severity) Handler {
	return NewWriteCloserHandler(os.Stdout, level)
}

// Instantiates a new Handler which will write on Stderr when level is reached.
func NewStderrHandler(level Severity) Handler {
	return NewWriteCloserHandler(os.Stderr, level)
}

func (wch *writeCloserHandler) Close() {
	wch.wc.Close()
}

func (wch *writeCloserHandler) Handle(r *Record) {
	wch.handler.Prepare(r)
	wch.wc.Write([]byte(r.Formatted))
}

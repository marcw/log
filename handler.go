// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

// A Handler is an interface that sends log to a destination
type Handler interface {
	IsHandling(Severity) bool // The handler accepts this severity
	Handle(*Record)           // Handle the log record
	PushProcessor(Processor)  // Push a new processor to the handler's stack
	PopProcessor()            // Removes a processor from the handler's stack
	SetFormatter(Formatter)   // Set the formatter for this handler
	GetFormatter() Formatter  // Returns the formatter used by this handler
	Close()                   // Close the handler
}

type handler struct {
	level      Severity
	formatter  Formatter
	processors []Processor
}

func (h *handler) IsHandling(level Severity) bool {
	return level <= h.level
}

func (h *handler) SetFormatter(f Formatter) {
	h.formatter = f
}

func (h *handler) GetFormatter() Formatter {
	return h.formatter
}

func (h *handler) Prepare(r *Record) {
	h.process(r)
	h.formatter.Format(r)
}

func (h *handler) process(r *Record) {
	for k := range h.processors {
		h.processors[k].Process(r)
	}
}

func (h *handler) PushProcessor(p Processor) {
	processors := make([]Processor, len(h.processors))

	copy(processors, h.processors)

	h.processors = []Processor{p}
	h.processors = append(h.processors, processors...)
}

func (h *handler) PopProcessor() {
	if len(h.processors) > 0 {
		h.processors = h.processors[1:len(h.processors)]
		return
	}

	panic("Processors stack is empty")
}

func (h *handler) Write() {
	// NO OP
}

// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

// HandlerInterface represents a type that sends log to a destination
type HandlerInterface interface {
	S(Severity) bool         // Returns true if the handler accepts this severity
	Handle(Record)           // Handle the log record
	PushProcessor(Processor) // Push a new processor to the handler's stack
	PopProcessor()           // Removes a processor from the handler's stack
	SetFormatter(Formatter)  // Set the formatter for this handler
	GetFormatter() Formatter // Returns the formatter used by this handler
	Close()                  // Close the handler
}

// Handler is a somewhat a "abstract" type you can embed in your own handler.
// It provides easiness in writing your own handlers
type Handler struct {
	Level      Severity
	Formatter  Formatter
	Processors []Processor
}

func (h *Handler) S(level Severity) bool {
	return level <= h.Level
}

func (h *Handler) SetFormatter(f Formatter) {
	h.Formatter = f
}

func (h *Handler) GetFormatter() Formatter {
	return h.Formatter
}

func (h *Handler) Prepare(r *Record) {
	h.Process(r)
	if h.Formatter == nil {
		h.Formatter = NewSimpleLineFormatter()
	}
	h.Formatter.Format(r)
}

func (h *Handler) Process(r *Record) {
	for k := range h.Processors {
		h.Processors[k].Process(r)
	}
}

func (h *Handler) PushProcessor(p Processor) {
	processors := make([]Processor, len(h.Processors))

	copy(processors, h.Processors)

	h.Processors = []Processor{p}
	h.Processors = append(h.Processors, processors...)
}

func (h *Handler) PopProcessor() {
	if len(h.Processors) > 0 {
		h.Processors = h.Processors[1:len(h.Processors)]
		return
	}

	panic("Processors stack is empty")
}

func (h *Handler) Write() {
	// NO OP
}

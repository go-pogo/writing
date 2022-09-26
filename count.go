// Copyright (c) 2022, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package writing

import (
	"io"
)

// CountingWriter counts the amount of bytes written to its Write method and
// collects any non-nil errors that occur.
type CountingWriter interface {
	io.Writer
	// Count returns the number of written bytes.
	Count() int
	// Errors returns the errors returned by any of the calls to Write.
	Errors() []error
	// Reset sets count to 0 and clears the errors slice.
	Reset()
}

// CountingStringWriter counts the amount of bytes written to its Write and
// WriteString methods and collects any non-nil errors that occur.
type CountingStringWriter interface {
	CountingWriter
	io.StringWriter
}

// ToCountingWriter wraps an io.Writer with a CountingWriter.
func ToCountingWriter(w io.Writer) CountingWriter {
	if cw, ok := w.(CountingWriter); ok {
		return cw
	}

	return &countingWriter{target: w}
}

// ToCountingStringWriter wraps an io.Writer with a CountingStringWriter.
func ToCountingStringWriter(w io.Writer) CountingStringWriter {
	if csw, ok := w.(CountingStringWriter); ok {
		return csw
	}

	res := &countingStringWriter{
		target: ToStringWriter(w),
	}
	if cw, ok := w.(*countingWriter); ok {
		res.countingWriter = cw
	} else {
		res.countingWriter = &countingWriter{target: w}
	}

	return res
}

type countingWriter struct {
	target io.Writer
	count  int
	errs   []error
}

func (cw *countingWriter) Count() int { return cw.count }

func (cw *countingWriter) Errors() []error { return cw.errs }

func (cw *countingWriter) Reset() {
	if r, ok := cw.target.(interface{ Reset() }); ok {
		r.Reset()
	}
	cw.count = 0
	cw.errs = cw.errs[:0]
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.target.Write(p)
	if err != nil {
		cw.errs = append(cw.errs, err)
	}

	cw.count += n
	return n, err
}

type countingStringWriter struct {
	*countingWriter
	target StringWriter
}

func (csw *countingStringWriter) WriteString(s string) (int, error) {
	n, err := csw.target.WriteString(s)
	if err != nil {
		csw.errs = append(csw.errs, err)
	}

	csw.count += n
	return n, err
}

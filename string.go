// Copyright (c) 2022, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package writing

import (
	"io"
)

// A StringWriter is both an io.Writer and an io.StringWriter.
type StringWriter interface {
	io.Writer
	io.StringWriter
}

// ToStringWriter wraps an io.Writer with a StringWriter.
func ToStringWriter(w io.Writer) StringWriter {
	if sw, ok := w.(StringWriter); ok {
		return sw
	}

	return &stringWriter{w}
}

type stringWriter struct{ io.Writer }

func (w *stringWriter) WriteString(s string) (int, error) {
	return w.Writer.Write([]byte(s))
}

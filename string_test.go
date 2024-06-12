// Copyright (c) 2023, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package writing

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToStringWriter(t *testing.T) {
	t.Run("same", func(t *testing.T) {
		var sb strings.Builder
		assert.Same(t, &sb, ToStringWriter(&sb))
	})
	t.Run("upgraded", func(t *testing.T) {
		var writer writerOnly
		sw := ToStringWriter(&writer)
		assert.NotSame(t, &writer, sw)
		_, _ = sw.WriteString("test")
		assert.Equal(t, "test", string(writer.buf))
	})
}

type writerOnly struct {
	buf []byte
}

func (w *writerOnly) Write(p []byte) (n int, err error) {
	if w.buf == nil {
		w.buf = make([]byte, 0, len(p))
	}

	w.buf = append(w.buf, p...)
	return len(p), nil
}

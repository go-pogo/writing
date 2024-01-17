// Copyright (c) 2023, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package writing

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestToStringWriter(t *testing.T) {
	t.Run("same", func(t *testing.T) {
		var sb strings.Builder
		assert.Same(t, &sb, ToStringWriter(&sb))
	})
	t.Run("upgraded", func(t *testing.T) {
		type Writer struct {
			io.Writer
		}

		var writer Writer
		assert.NotSame(t, &writer, ToStringWriter(&writer))
	})
}

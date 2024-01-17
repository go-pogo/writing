// Copyright (c) 2023, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package writing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytesBufferPool(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		var p BytesBufferPool
		assert.Equal(t, 0, p.Get().Cap())
	})
	t.Run("NewBytesBufferPool", func(t *testing.T) {
		const n = 12
		p := NewBytesBufferPool(n)
		assert.True(t, p.Get().Cap() >= n)
	})
}

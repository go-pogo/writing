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

func TestBytesBufferPool_Get(t *testing.T) {
	var p BytesBufferPool

	b := p.Get()
	assert.Equal(t, 0, b.Cap())

	const str = "dit is een hele lang test string met heel veel willekeurige woorden die nergens op slaan zolang de buffer maar groeit"
	b.WriteString(str)
	assert.Equal(t, str, b.String())

	p.Put(b)
	b = p.Get()
	assert.Equal(t, 0, b.Len())
	assert.Equal(t, 128, b.Cap())
}

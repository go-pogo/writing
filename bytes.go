// Copyright (c) 2022, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package writing

import (
	"bytes"
	"sync"
)

// A BytesBufferPool is a sync.Pool which stores bytes.Buffer(s).
type BytesBufferPool struct {
	p sync.Pool
}

// NewBytesBufferPool creates a new BytesBufferPool with a given initial
// capacity n for each retrieved bytes.Buffer.
func NewBytesBufferPool(n int) *BytesBufferPool {
	return &BytesBufferPool{p: sync.Pool{
		New: func() interface{} {
			var buf bytes.Buffer
			if n != 0 {
				buf.Grow(n)
			}
			return &buf
		},
	}}
}

// Get selects an arbitrary bytes.Buffer from the BytesBufferPool, removes it
// from the pool, and returns it to the caller.
func (p *BytesBufferPool) Get() *bytes.Buffer {
	return p.p.Get().(*bytes.Buffer)
}

// Put adds the bytes.Buffer b to the pool.
func (p *BytesBufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	p.p.Put(b)
}

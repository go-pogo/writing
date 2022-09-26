// Copyright (c) 2022, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package writing

import (
	"bytes"
	"sync"
)

type BytesBufferPool struct {
	p sync.Pool
}

func NewBytesBufferPool(n int) *BytesBufferPool {
	return &BytesBufferPool{p: sync.Pool{
		New: func() interface{} {
			var buf bytes.Buffer
			buf.Grow(n)
			return &buf
		},
	}}
}

func (p *BytesBufferPool) Get() *bytes.Buffer {
	return p.p.Get().(*bytes.Buffer)
}

func (p *BytesBufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	p.p.Put(b)
}

// Copyright (c) 2024, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package writing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCountingWriter(t *testing.T) {
	var writer writerOnly
	csw := ToCountingWriter(&writer)
	assert.NotSame(t, &writer, csw)

	want := []byte("test")
	haveCount, haveErr := csw.Write(want)

	assert.Equal(t, want, writer.buf)
	assert.Equal(t, 4, haveCount)
	assert.Equal(t, haveCount, csw.Count())
	assert.Nil(t, haveErr)
}

func TestToCountingStringWriter(t *testing.T) {
	var writer writerOnly
	csw := ToCountingStringWriter(&writer)
	assert.NotSame(t, &writer, csw)

	const want = "test"
	haveCount, haveErr := csw.WriteString(want)

	assert.Equal(t, want, string(writer.buf))
	assert.Equal(t, 4, haveCount)
	assert.Equal(t, haveCount, csw.Count())
	assert.Nil(t, haveErr)
}

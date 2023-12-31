// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dig

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionStrings(t *testing.T) {
	t.Parallel()

	t.Run("DeferAcyclicVerification", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "DeferAcyclicVerification()", fmt.Sprint(DeferAcyclicVerification()))
	})

	t.Run("setRand", func(t *testing.T) {
		t.Parallel()

		t.Run("nil", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "setRand(0x0)", fmt.Sprint(setRand(nil)))
		})

		t.Run("non nil", func(t *testing.T) {
			t.Parallel()

			opt := setRand(rand.New(rand.NewSource(42)))
			assert.NotEqual(t, "setRand(0x0)", fmt.Sprint(opt))
			assert.Contains(t, fmt.Sprint(opt), "setRand(0x")
		})
	})

	t.Run("DryRun", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "DryRun(true)", fmt.Sprint(DryRun(true)))
		assert.Equal(t, "DryRun(false)", fmt.Sprint(DryRun(false)))
	})

	t.Run("RecoverFromPanics()", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "RecoverFromPanics()", fmt.Sprint(RecoverFromPanics()))
	})
}

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package embeddirective

import (
	"io"
	"os"
)

//go:embed embed.txt //@suggestedfix("//go:embed", "quickfix", "")
var t string

func unused() {
	_ = os.Stdin
	_ = io.EOF
}

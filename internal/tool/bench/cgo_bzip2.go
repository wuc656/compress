// Copyright 2015, Joe Tsai. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

//go:build cgo && !no_cgo_bzip2
// +build cgo,!no_cgo_bzip2

package main

import "github.com/wuc656/compress/internal/cgo/bzip2"

func init() {
	RegisterEncoder(FormatBZ2, "cgo", bzip2.NewWriter)
	RegisterDecoder(FormatBZ2, "cgo", bzip2.NewReader)
}

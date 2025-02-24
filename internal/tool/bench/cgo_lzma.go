// Copyright 2016, Joe Tsai. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

//go:build cgo && !no_cgo_lzma
// +build cgo,!no_cgo_lzma

package main

import "github.com/wuc656/compress/internal/cgo/lzma"

func init() {
	RegisterEncoder(FormatLZMA2, "cgo", lzma.NewWriter)
	RegisterDecoder(FormatLZMA2, "cgo", lzma.NewReader)
}

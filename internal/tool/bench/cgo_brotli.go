// Copyright 2016, Joe Tsai. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

//go:build cgo && !no_cgo_brotli
// +build cgo,!no_cgo_brotli

package main

import "github.com/wuc656/compress/internal/cgo/brotli"

func init() {
	RegisterEncoder(FormatBrotli, "cgo", brotli.NewWriter)
	RegisterDecoder(FormatBrotli, "cgo", brotli.NewReader)
}

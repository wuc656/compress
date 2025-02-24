// Copyright 2016, Joe Tsai. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

//go:build cgo && !no_cgo_zstd
// +build cgo,!no_cgo_zstd

package main

import "github.com/wuc656/compress/internal/cgo/zstd"

func init() {
	RegisterEncoder(FormatZstd, "cgo", zstd.NewWriter)
	RegisterDecoder(FormatZstd, "cgo", zstd.NewReader)
}

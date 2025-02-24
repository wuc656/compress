// Copyright 2015, Joe Tsai. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

//go:build cgo && !no_cgo_flate
// +build cgo,!no_cgo_flate

package main

import "github.com/wuc656/compress/internal/cgo/flate"

func init() {
	RegisterEncoder(FormatFlate, "cgo", flate.NewWriter)
	RegisterDecoder(FormatFlate, "cgo", flate.NewReader)
}

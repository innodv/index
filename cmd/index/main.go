/**
 * Copyright 2019 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"

	"github.com/innodv/index"
)

func main() {
	entries, err := index.GetEntries(10000)
	if err != nil {
		panic(err)
	}
	fmt.Println(entries)
}

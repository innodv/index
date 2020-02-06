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
	out := map[string]interface{}{}
	target := 100000
	entries := index.ConsumeEntries()
	cnt := 0
	for entry := range entries {
		if cnt == target {
			break
		}
		name := entry["Path"] + entry["Version"]
		out[name] = nil
		cnt++
		if cnt%10000 == 0 {
			fmt.Printf("%d\n", cnt)
		}
	}
	if cnt != target {
		fmt.Println("failed")
	}
	fmt.Println("success")
}

/**
 * Copyright 2019 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package index

import (
	"bufio"
	"encoding/json"
	"net/http"
)

func GetEntries(cnt int) ([]map[string]string, error) {
	out := []map[string]string{}
	timestamp := ""
	for i := 0; i < cnt; i += 2000 {
		resp, err := http.Get("https://index.golang.org/index?since=" + timestamp)
		if err != nil {
			return nil, err
		}
		rdr := bufio.NewReader(resp.Body)
		for line, err := rdr.ReadBytes('\n'); err == nil; line, err = rdr.ReadBytes('\n') {
			var data map[string]string
			err := json.Unmarshal(line, &data)
			if err != nil {
				return nil, err
			}
			out = append(out, data)
			timestamp = data["Timestamp"]
		}
	}
	return out, nil
}

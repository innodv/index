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
	"time"
)

func GetEntries(timestamp string) ([]map[string]string, error) {
	out := []map[string]string{}
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
	}
	return out, nil
}

func ConsumeEntries() <-chan map[string]string {
	out := make(chan map[string]string)
	timestamp := ""
	go func() {
		for i := 0; true; i += 2000 {
			ents, err := GetEntries(timestamp)
			if err != nil {
				time.Sleep(5 * time.Minute)
				continue
			}
			for i := range ents {
				out <- ents[i]
			}
			timestamp = ents[len(ents)-1]["Timestamp"]
		}
	}()
	return out
}

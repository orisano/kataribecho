// Copyright 2020-2021 Nao Yonashiro
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type LogEntry struct {
	Method   string `json:"method"`
	URI      string `json:"uri"`
	Status   int    `json:"status"`
	Latency  int    `json:"latency"`
	BytesOut int    `json:"bytes_out"`
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var e LogEntry
		err := json.Unmarshal(s.Bytes(), &e)
		if err != nil {
			continue
		}
		if e.Method == "" || e.Status == 0 {
			continue
		}
		fmt.Printf("%q %v %v %f\n", e.Method+" "+e.URI, e.Status, e.BytesOut, time.Duration(e.Latency).Seconds())
	}
}

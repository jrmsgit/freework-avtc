// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package record

import (
	"fmt"
	"time"
)

// $ date -d 'TZ="UTC" 20200220 00:00:00' '+%s'
// 1582156800
var tsStart int64 = 1582156800

// Gen populate log file with random data
func Gen() {
	tsEnd := time.Now().UTC().Unix()
	for sec := tsStart; sec <= tsEnd; sec++ {
		fmt.Println(sec)
	}
}

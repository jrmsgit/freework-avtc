// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package record

import (
	"fmt"
	"math/rand"
	"time"
)

// $ date -d 'TZ="UTC" 20200220 00:00:00' '+%s'
// 1582156800
var tsStart int64 = 1582156800

var logfmt = `{"level":"debug","ts":%d,"conn_id":%d,"state":"closed","Tx":%d,"Rx":%d}`
var xmax = 10000

// Gen populate log file with random data
func Gen() {
	tsEnd := time.Now().UTC().Unix()
	rand.Seed(tsEnd)

	//~ tsStart = tsEnd - 2 // FIXME

	for x := tsStart; x <= tsEnd; x++ {
		fmt.Printf(logfmt + "\n", tsStart + x, x, rand.Intn(xmax), rand.Intn(xmax))
	}
}

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

// generate log data, if p is true print it to stdout too
func genlog(p bool) {
	tsEnd := time.Now().UTC().Unix()
	rand.Seed(tsEnd)

	//~ tsStart = tsEnd - 2 // FIXME

	for x := tsStart; x <= tsEnd; x++ {
		// tx and rx are generated here so we don't cheat the benchmarks
		tx := rand.Intn(xmax)
		rx := rand.Intn(xmax)
		if p {
			fmt.Printf(logfmt + "\n", tsStart + x, x, tx, rx)
		}
	}
}

// Gen generates log data but do not prints it.
func Gen() {
	genlog(false)
}

// Print generates log data and prints it to stdout.
func Print() {
	genlog(true)
}

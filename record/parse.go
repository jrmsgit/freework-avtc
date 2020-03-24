// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package record

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jrmsgit/freework-avtc/hist/rx"
	"github.com/jrmsgit/freework-avtc/hist/tx"
)

// Parse parses log lines data and generates prometheus histograms (tx and rx)
func Parse() {
	txHist := tx.Hist()
	rxHist := rx.Hist()

	//~ tx.Hist()
	//~ rx.Hist()

	file, err := os.Open("./out/log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		r := New(scanner.Bytes())

		//~ fmt.Printf("%d Tx: %d Rx: %d\n", r.ConnId, r.Tx, r.Rx)

		txHist.WithLabelValues(fmt.Sprintf("%d", r.ConnId)).Observe(float64(r.Tx))
		rxHist.WithLabelValues(fmt.Sprintf("%d", r.ConnId)).Observe(float64(r.Rx))
	}
}

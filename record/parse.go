// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package record

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jrmsgit/freework-avtc/hist/rx"
	"github.com/jrmsgit/freework-avtc/hist/tx"
)

// Parse parses log lines data and generates prometheus histograms (tx and rx)
func Parse() {
	txHist := tx.Hist()
	rxHist := rx.Hist()

	file, err := os.Open("./out/log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var r Record

	for scanner.Scan() {

		err := json.Unmarshal(scanner.Bytes(), &r)
		if err != nil {
			panic(err)
		}

		txHist.WithLabelValues(fmt.Sprintf("%d", r.ConnId)).Observe(float64(r.Tx))
		rxHist.WithLabelValues(fmt.Sprintf("%d", r.ConnId)).Observe(float64(r.Rx))
	}
}

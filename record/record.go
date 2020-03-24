// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package record

import (
	"encoding/json"
)

type Record struct {
	Level  string
	TStamp int64
	ConnId int64 `json:"conn_id"`
	State  string
	Tx     int
	Rx     int
}

func New(blob []byte) *Record {
	var r Record
	err := json.Unmarshal(blob, &r)
	if err != nil {
		panic(err)
	}
	return &r
}

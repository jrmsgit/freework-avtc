// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package record

type Record struct {
	Level  string `json:",omit"`
	TStamp int64 `json:",omit"`
	ConnId int64 `json:"conn_id"`
	State  string `json:",omit"`
	Tx     int
	Rx     int
}

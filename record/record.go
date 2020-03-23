// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package record

//~ import (
//~ "time"
//~ )

type Record struct {
	Level  string
	TStamp int64
	ConnID int64
	State  string
	Tx     int
	Rx     int
}

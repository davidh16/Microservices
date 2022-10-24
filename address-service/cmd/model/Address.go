package model

import "time"

type Address struct {
	Id       int64
	Location string
	Valid    bool
	ValidAt  time.Time
}

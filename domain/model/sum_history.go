package model

import "time"

type SumHistory struct {
	Date time.Time
	Numbers Numbers
	Result int
}

package model

import "time"

type SumHistory struct {
	IsNorm  bool
	Date    time.Time
	Numbers Numbers
	Result  int
}

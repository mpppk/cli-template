package repository

import "github.com/mpppk/cli-template/domain/model"

type SumHistory interface {
	AddHistory(sumHistory model.SumHistory)
}

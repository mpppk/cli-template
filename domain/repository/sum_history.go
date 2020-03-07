package repository

import "github.com/mpppk/cli-template/domain/model"

type SumHistory interface {
	Add(sumHistory *model.SumHistory)
	List(limit int) []*model.SumHistory
}

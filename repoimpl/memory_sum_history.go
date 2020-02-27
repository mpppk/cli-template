package repoimpl

import (
	"github.com/mpppk/cli-template/domain/model"
	"github.com/mpppk/cli-template/domain/repository"
)

type MemorySumHistory struct {
	history []model.SumHistory
}

func NewMemorySumHistory() repository.SumHistory {
	return &MemorySumHistory{}
}

func (m *MemorySumHistory) AddHistory(sumHistory model.SumHistory) {
	m.history = append(m.history, sumHistory)
}


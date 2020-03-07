package repoimpl

import (
	"log"

	"github.com/mpppk/cli-template/domain/model"
	"github.com/mpppk/cli-template/domain/repository"
)

type MemorySumHistory struct {
	history []*model.SumHistory
}

func NewMemorySumHistory() repository.SumHistory {
	return &MemorySumHistory{history: []*model.SumHistory{}}
}

func (m *MemorySumHistory) Add(sumHistory *model.SumHistory) {
	m.history = append(m.history, sumHistory)
	log.Printf("current history: %v", m.history)
}

func (m *MemorySumHistory) List(limit int) []*model.SumHistory {
	l := len(m.history)
	if l <= limit {
		return m.history
	}
	return m.history[l-limit : l]
}

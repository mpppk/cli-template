package usecase

import (
	"github.com/mpppk/cli-template/domain/model"
	"github.com/mpppk/cli-template/domain/repository"
	"log"
	"time"
)

type SumUseCase struct {
	sumHistoryRepository repository.SumHistory
}

func NewSum(sumHistoryRepository repository.SumHistory) *SumUseCase {
	return &SumUseCase{
		sumHistoryRepository: sumHistoryRepository,
	}
}

// CalcSum is use case to calculate sum
func (s *SumUseCase) CalcSum(numbers []int) int {
	result := model.NewNumbers(numbers).CalcSum()
	s.sumHistoryRepository.AddHistory(model.SumHistory{
		Date:   time.Now(), // FIXME
		Numbers: numbers,
		Result: result,
	})
	return result
}

// CalcL1Norm is use case to calculate L1 norm
func (s *SumUseCase) CalcL1Norm(numbers []int) int {
	result := model.NewNumbers(numbers).CalcL1Norm()
	now := time.Now()
	log.Printf("start saving history. date=%v, numbers=%q, result=%q\n", now, numbers, result)
	s.sumHistoryRepository.AddHistory(model.SumHistory{
		Date:   now, // FIXME
		Numbers: numbers,
		Result: result,
	})
	log.Printf("finish saving history. date=%v, numbers=%q, result=%q\n", now, numbers, result)
	return result
}

package service

import (
	"errors"
	"go-math-api/dto"
	"go-math-api/repository"
)

type CalculatorService interface {
	Add(request *dto.CalculatorRequest) int
	Substract(request *dto.CalculatorRequest) int
	Multiply(request *dto.CalculatorRequest) int
	Divide(request *dto.CalculatorRequest) (int, error)
}

type calculatorService struct {
	repository repository.CalculatorRepository
}

func NewCalculatorService(r repository.CalculatorRepository) *calculatorService {
	return &calculatorService{r}
}

func (s *calculatorService) Add(request *dto.CalculatorRequest) int {
	return request.Num1 + request.Num2
}

func (s *calculatorService) Substract(request *dto.CalculatorRequest) int {
	return request.Num1 - request.Num2
}

func (s *calculatorService) Multiply(request *dto.CalculatorRequest) int {
	return request.Num1 * request.Num2
}

func (s *calculatorService) Divide(request *dto.CalculatorRequest) (int, error) {
	if request.Num2 == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return request.Num1 / request.Num2, nil
}

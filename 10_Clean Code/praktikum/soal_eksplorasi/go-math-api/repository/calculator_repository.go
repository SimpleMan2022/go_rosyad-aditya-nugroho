package repository

type CalculatorRepository interface {
}
type calculatorRepository struct {
}

func NewCalculatorRepository() *calculatorRepository {
	return &calculatorRepository{}
}

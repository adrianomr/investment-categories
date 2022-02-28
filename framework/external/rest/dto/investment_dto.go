package dto

type InvestmentDto struct {
	ID            string
	Name          string
	Grade         float32
	Origin        string
	CurrentAmount float32
	TargetAmount  float32
	Category      *CategoryDto
	Investments   *[]InvestmentDto
}

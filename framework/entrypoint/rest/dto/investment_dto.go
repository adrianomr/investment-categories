package dto

type InvestmentDto struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	Grade         float32      `json:"grade"`
	Origin        string       `json:"origin"`
	CurrentAmount float32      `json:"currentAmount"`
	TargetAmount  float32      `json:"targetAmount"`
	Category      *CategoryDto `json:"category"`
}

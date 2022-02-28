package dto

type CategoryDto struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	Grade         float32      `json:"grade"`
	CurrentAmount float32      `json:"currentAmount"`
	TargetAmount  float32      `json:"TargetAmount"`
	Category      *CategoryDto `json:"category"`
}

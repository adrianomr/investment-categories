package dto

type Investment struct {
	ID           string
	Name         string
	Grade        float32
	Origin       string
	CurrentAmount float32
	TargetAmount float32
	Category    *CategoryDto
}

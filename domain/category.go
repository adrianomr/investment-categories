package domain

type Category struct {
	ID             string
	Name           string
	Grade          float32
	CurrentAmount  float32
	InvestedAmount float32
	TargetAmount   float32
	Balance        float32
	Category       *Category
	Investments    []*Investment
	UserId         int
}

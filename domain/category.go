package domain

type Category struct {
	ID            string
	Name          string
	Grade         float32
	CurrentAmount float32
	TargetAmount  float32
	Category      *Category
	UserId        int
}

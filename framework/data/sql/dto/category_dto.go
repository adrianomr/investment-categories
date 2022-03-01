package dto

type CategoryDto struct {
	ID            string
	Name          string
	Grade         float32
	CurrentAmount float32
	TargetAmount  float32
	UserId        int
	CategoryId    string
	Category      *CategoryDto  `sql:"ForeignKey:CategoryId"`
	Investments   *[]Investment `sql:"ForeignKey:CategoryId"`
}

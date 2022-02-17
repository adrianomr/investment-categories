package dto

type CategoryDto struct {
	ID            string
	Name          string
	Grade         float32
	CurrentAmount float32
	TargetAmount  float32
	UserId        int32
	CategoryId    string
	Category      *CategoryDto  `gorm:"ForeignKey:CategoryId"`
	Investments   *[]Investment `gorm:"ForeignKey:CategoryId"`
}

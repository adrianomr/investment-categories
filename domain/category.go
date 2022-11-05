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

func (c *Category) CalculateTarget(categories *[]Category) {
	totalAmount := c.CurrentAmount
	sumGrade := c.Grade
	for _, category := range *categories {
		if c.ID != category.ID {
			totalAmount += category.CurrentAmount
			sumGrade += category.Grade
		}
	}
	c.TargetAmount = (c.Grade / sumGrade) * totalAmount
}

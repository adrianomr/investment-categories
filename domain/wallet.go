package domain

type Wallet struct {
	TotalAmount       float32     `json:"TotalAmount"`
	InvestedAmount    float32     `json:"investedAmount"`
	Balance           float32     `json:"balance"`
	PercentageBalance float32     `json:"percentageBalance"`
	Categories        *[]Category `json:"category"`
}

func (wallet *Wallet) Calculate() {
	for _, category := range *wallet.Categories {
		wallet.TotalAmount += category.CurrentAmount
		wallet.InvestedAmount += category.InvestedAmount
		wallet.Balance += category.Balance
		wallet.PercentageBalance = ((wallet.TotalAmount - wallet.InvestedAmount) / wallet.InvestedAmount) * 100
	}
}

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
		if wallet.Balance > 0 && wallet.InvestedAmount > 0 {
			wallet.PercentageBalance = (wallet.Balance / wallet.InvestedAmount) * 100
		}
	}
}

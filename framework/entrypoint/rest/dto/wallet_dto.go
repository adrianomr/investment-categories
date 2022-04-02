package dto

type WalletDto struct {
	TotalAmount       float32        `json:"TotalAmount"`
	InvestedAmount    float32        `json:"investedAmount"`
	Balance           float32        `json:"balance"`
	PercentageBalance float32        `json:"percentageBalance"`
	Categories        *[]CategoryDto `json:"category"`
}

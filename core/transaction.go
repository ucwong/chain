package core

import (
//	"strings"
)

// Transaction ...
type Transaction struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    uint64 `json:"amount"`
}

func (tx Transaction) Hash() string {
	return ""
}

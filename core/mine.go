package core

type Mine struct {
	Message      string        `json:"message"`
	Index        uint64        `json:"index"`
	Transactions []Transaction `json:"transactions"`
	Proof        uint64        `json:"proof"`
	PreviousHash string        `json:"previous_hash"`
}

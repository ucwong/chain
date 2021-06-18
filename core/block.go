package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Block struct {
	Index        uint64        `json:"index"`
	Timestamp    time.Time     `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Proof        uint64        `json:"proof"`
	PreviousHash string        `json:"previous_hash"`
}

func (b Block) Hash() string {
	jsonBytes, _ := json.Marshal(b)
	hashBytes := sha256.Sum256(jsonBytes)
	return hex.EncodeToString(hashBytes[:])
}

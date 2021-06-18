package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var (
	reverso string = "hello"
)

func Genesis() string {
	hash := sha256.Sum256([]byte(reverso))
	return hex.EncodeToString(hash[:])
}

func ValidProof(lastProof, proof uint64, lastHash string) bool {
	guess := fmt.Sprintf("%x%x%x", lastProof, proof, lastHash)
	guessBytes := sha256.Sum256([]byte(guess))
	guessHash := hex.EncodeToString(guessBytes[:])
	return guessHash[:4] == "0000"
}

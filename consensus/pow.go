package consensus

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func ValidProof(lastProof, proof uint64, lastHash string) bool {
	guess := fmt.Sprintf("%x%x%x", lastProof, proof, lastHash)
	guessBytes := sha256.Sum256([]byte(guess))
	guessHash := hex.EncodeToString(guessBytes[:])

	// constant diff
	return guessHash[:4] == "0000"
}

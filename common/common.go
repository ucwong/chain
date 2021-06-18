package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var (
//reverso string = "hello"
//2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
)

func Genesis() string {
	return "2494bd63182a60fa55980962d812364d88d56585a713efa80e2c81abdd492fc2"
}

func ValidProof(lastProof, proof uint64, lastHash string) bool {
	guess := fmt.Sprintf("%x%x%x", lastProof, proof, lastHash)
	guessBytes := sha256.Sum256([]byte(guess))
	guessHash := hex.EncodeToString(guessBytes[:])

	if guessHash[:4] == "0000" {
		fmt.Println(lastProof)
		fmt.Println(proof)
		fmt.Println(lastHash)

		fmt.Println(guessHash)
		return true
	}

	return false
}

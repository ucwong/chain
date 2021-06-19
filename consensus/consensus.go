package consensus

import (
	"github.com/ucwong/chain/core/types"
	"log"
)

var (
//reverso string = "hello"
//2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
)

func Genesis() string {
	return "2494bd63182a60fa55980962d812364d88d56585a713efa80e2c81abdd492fc2"
}

func ValidChain(chain []types.Block) bool {
	lastBlock := chain[0]
	currentIndex := 1
	for currentIndex < len(chain) {
		block := chain[currentIndex]
		if block.PreviousHash != lastBlock.Hash() {
			log.Printf("Invalid hash %s, %s\n", block.PreviousHash, lastBlock.Hash())
			return false
		}
		if !ValidProof(lastBlock.Proof, block.Proof, lastBlock.Hash()) {
			log.Printf("Invalid proof %d, %d %s\n", lastBlock.Proof, block.Proof, lastBlock.PreviousHash)
			return false
		}
		lastBlock = block
		currentIndex++
	}
	return true
}

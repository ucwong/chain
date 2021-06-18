package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

type Blockchain struct {
	Chain               []Block       `json:"chain"`
	CurrentTransactions []Transaction `json:"current_transactions"`
	Nodes               []string      `json:"nodes"`
	nodeIdentifier      string
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CurrentTransactions = nil
	bc.Chain = nil
	bc.Nodes = nil
	bc.newBlock(uint64(21), Genesis())
	out, _ := exec.Command("uuidgen").Output()
	bc.nodeIdentifier = strings.Replace(string(out), "-", "", -1)

	return bc
}

func (bc *Blockchain) registerNode(address string) {
	parsedURL, err := url.Parse(address)
	if err != nil {
		log.Println(err)
		return
	}
	host := parsedURL.Host
	if host == "" {
		return
	}
	for _, node := range bc.Nodes {
		if node == host {
			return
		}
	}
	bc.Nodes = append(bc.Nodes, host)
}

func (bc *Blockchain) validChain(chain []Block) bool {
	lastBlock := chain[0]
	currentIndex := 1
	for currentIndex < len(chain) {
		block := chain[currentIndex]
		if block.PreviousHash != lastBlock.Hash() {
			return false
		}
		if !ValidProof(lastBlock.Proof, block.Proof, lastBlock.PreviousHash) {
			return false
		}
		lastBlock = block
		currentIndex++
	}
	return true
}

func (bc *Blockchain) resolveConflicts() bool {
	var newChain []Block
	neighbours := bc.Nodes
	maxLength := len(bc.Chain)
	for _, node := range neighbours {
		url := fmt.Sprintf("http://%s.chain", node)
		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
			return false
		}
		defer res.Body.Close()
		byteArr, _ := ioutil.ReadAll(res.Body)
		var response chainResponse
		if err = json.Unmarshal(byteArr, &response); err != nil {
			log.Println(err)
			return false
		}
		length := response.Length
		chain := response.Chain
		if length > maxLength && bc.validChain(chain) {
			maxLength = length
			newChain = chain
		}
	}
	if len(newChain) > 0 {
		bc.Chain = newChain
		return true
	}
	return false

}

func (bc *Blockchain) newBlock(proof uint64, previousHash string) Block {
	block := Block{
		Index:        uint64(len(bc.Chain) + 1),
		Timestamp:    time.Now(),
		Transactions: bc.CurrentTransactions,
		Proof:        proof,
		PreviousHash: previousHash,
	}
	bc.CurrentTransactions = nil
	bc.Chain = append(bc.Chain, block)
	return block
}

func (bc *Blockchain) newTransaction(sender, recipient string, amount uint64) uint64 {
	transaction := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
	bc.CurrentTransactions = append(bc.CurrentTransactions, transaction)
	return bc.lastBlock().Index + 1
}

func (bc *Blockchain) proofOfWork(lastBlock Block) uint64 {
	lastProof := lastBlock.Proof
	lastHash := lastBlock.Hash()
	proof := uint64(0)
	for ValidProof(lastProof, proof, lastHash) == false {
		proof++
	}
	return proof
}

func (bc *Blockchain) lastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}
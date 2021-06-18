package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type chainResponse struct {
	Length int     `json:"length"`
	Chain  []Block `json:"chain"`
}

func (bc *Blockchain) MineHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
	lastBlock := bc.lastBlock()
	proof := bc.proofOfWork(lastBlock)

	bc.newTransaction("0", bc.nodeIdentifier, 1)
	previousHash := lastBlock.Hash()
	block := bc.newBlock(proof, previousHash)

	mine := Mine{
		Message:      "New Block Forged",
		Index:        block.Index,
		Transactions: block.Transactions,
		Proof:        block.Proof,
		PreviousHash: block.PreviousHash,
	}
	res, err := json.Marshal(mine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (bc *Blockchain) NewTransactionHandler(w http.ResponseWriter, r *http.Request) {
	type jsonBody struct {
		Sender    string `json:"sender"`
		Recipient string `json:"recipient"`
		Amount    uint64 `json:"amount"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}

	index := bc.newTransaction(b.Sender, b.Recipient, b.Amount)
	type response struct {
		Message string `json:"message"`
	}
	resNewTransaction := response{
		Message: "Transaction will be added to Block" + strconv.FormatUint(index, 10),
	}
	res, err := json.Marshal(resNewTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (bc *Blockchain) ChainHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Chain  []Block `json:"chain"`
		Length int     `json:"length"`
	}
	resChain := response{
		Chain:  bc.Chain,
		Length: len(bc.Chain),
	}
	res, err := json.Marshal(resChain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (bc *Blockchain) NodesRegisterHandler(w http.ResponseWriter, r *http.Request) {
	type jsonBody struct {
		Nodes []string `json:"nodes"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, node := range b.Nodes {
		bc.registerNode(node)
	}
	type response struct {
		Message    string   `json:"message"`
		TotalNodes []string `json:"total_nodes"`
	}
	var resNodesRegister response
	resNodesRegister = response{
		Message:    "New nodes have been added",
		TotalNodes: bc.Nodes,
	}
	res, err := json.Marshal(resNodesRegister)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (bc *Blockchain) NodesResolveHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Message string  `json:"message"`
		Chain   []Block `json:"chain"`
	}
	var resNodesResolve response
	replaced := bc.resolveConflicts()
	if replaced {
		resNodesResolve = response{
			Message: "Our chain was replaced",
			Chain:   bc.Chain,
		}
	} else {
		resNodesResolve = response{
			Message: "Our chain is authoritative",
			Chain:   bc.Chain,
		}
	}
	res, err := json.Marshal(resNodesResolve)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

package main

import (
	"fmt"
	"github.com/ucwong/chain/core"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var bc *core.Blockchain
var nodeIdentifier string

func main() {
	args := os.Args
	port := "5000"
	if len(args) > 2 && args[1] == "-p" {
		port = args[2]
	}
	bc = core.NewBlockchain()
	out, _ := exec.Command("uuidgen").Output()
	nodeIdentifier = strings.Replace(string(out), "-", "", -1)

	http.HandleFunc("/mine", bc.MineHandler)
	http.HandleFunc("/transactions/new", bc.NewTransactionHandler)
	http.HandleFunc("/chain", bc.ChainHandler)
	http.HandleFunc("/nodes/register", bc.NodesRegisterHandler)
	http.HandleFunc("/nodes/resolve", bc.NodesResolveHandler)

	log.Printf("Server listening on localhost:%s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}

}

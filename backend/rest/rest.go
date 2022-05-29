package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minjaelee0727/idWallet/backend/blockchain"
	"github.com/minjaelee0727/idWallet/backend/constant"
	"github.com/minjaelee0727/idWallet/backend/utils"
	"github.com/minjaelee0727/idWallet/backend/wallet"
)

const (
	port string = ":4000"
)

type walletResponse struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"secretKey"`
}

func createIdWallet(rw http.ResponseWriter, r *http.Request) {
	var credential constant.Credential
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&credential))
	w := wallet.CreateWallet()
	blockchain.Blockchain().AddBlock(credential, w)
	json.NewEncoder(rw).Encode(walletResponse{PublicKey: w.PublicKey, PrivateKey: w.PrivateKey})
}

func verifySignature(rw http.ResponseWriter, r *http.Request) {

}

func verifyAge(rw http.ResponseWriter, r *http.Request) {

}

func seeBlocks(rw http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blocks(blockchain.Blockchain())))
}

///

func StartService() {
	router := mux.NewRouter()
	router.Use(utils.JsonContentTypeMiddleware, utils.LoggerMiddleware)
	router.HandleFunc("/create", createIdWallet).Methods("POST")
	router.HandleFunc("/verify", verifySignature).Methods("GET")
	router.HandleFunc("/blocks", seeBlocks).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

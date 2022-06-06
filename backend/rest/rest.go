package rest

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minjaelee0727/idWallet/backend/blockchain"
	"github.com/minjaelee0727/idWallet/backend/constant"
	"github.com/minjaelee0727/idWallet/backend/utils"
	"github.com/minjaelee0727/idWallet/backend/wallet"
)

type walletResponse struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"secretKey"`
}

type rangeProofData struct {
	t1 string
	t2 string
	t3 string
	t4 string
	td string
	s  string
	z  string
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

// func verifyAge(rw http.ResponseWriter, r *http.Request) {
// 	var input rangeProofData
// 	utils.HandleErr(json.NewDecoder(r.Body).Decode(&input))
// 	t1 := input.s1 * input.z ^ input.r1

// }

type SinProofData struct {
	Y string
	T string
	R string
	C string
}

type sinVerifyResult struct {
	Result bool `json:"result"`
}

type randomGResponse struct {
	G string
}

func verifySin(rw http.ResponseWriter, r *http.Request) {
	var spd SinProofData
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&spd))
	t, _ := new(big.Int).SetString(spd.T, 10)
	rR, _ := new(big.Int).SetString(spd.R, 10)
	y, _ := new(big.Int).SetString(spd.Y, 10)
	c, _ := new(big.Int).SetString(spd.C, 10)
	var a, b, d big.Int
	utils.HandleErr(json.NewEncoder(rw).Encode(sinVerifyResult{Result: t == d.Mul(a.Exp(gR, rR, nil), b.Exp(y, c, nil))}))
}

var gR *big.Int

func requestG(rw http.ResponseWriter, r *http.Request) {
	gR, _ = rand.Prime(rand.Reader, 20)
	utils.HandleErr(json.NewEncoder(rw).Encode(randomGResponse{G: gR.String()}))
}

func seeBlocks(rw http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blocks(blockchain.Blockchain())))
}

///

func Start(port int) {
	router := mux.NewRouter()
	router.Use(utils.JsonContentTypeMiddleware, utils.LoggerMiddleware)
	router.HandleFunc("/register", createIdWallet).Methods("POST")
	router.HandleFunc("/verify/sin", verifySin).Methods("GET")
	router.HandleFunc("/blocks", seeBlocks).Methods("GET")
	router.HandleFunc("/request/g", requestG).Methods("POST")
	fmt.Printf("Listening on http://0.0.0.0:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

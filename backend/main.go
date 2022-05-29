package main

import (
	"github.com/minjaelee0727/idWallet/backend/db"
	"github.com/minjaelee0727/idWallet/backend/rest"
)

func main() {
	defer db.CloseDB()
	db.DB()
	rest.StartService()
}

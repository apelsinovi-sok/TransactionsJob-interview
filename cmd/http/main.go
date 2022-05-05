package main

import (
	"test/infrastructure/DB"
	"test/internal/dilivery/http"
	"test/internal/repository"
)

func main() {
	db := DB.New()
	transferRep := repository.NewTransferRep(db)
	server := http.New(transferRep)
	server.Run()
}

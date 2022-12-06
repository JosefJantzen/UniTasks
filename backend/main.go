package main

import (
	"fmt"

	_ "github.com/cockroachdb/cockroach-go/v2/crdb"

	"unitasks.josefjantzen.de/backend/api"
	"unitasks.josefjantzen.de/backend/database"
)

func main() {
	dbService := database.InitDB()
	api.InitServer(dbService)
	fmt.Println("Completly Started")
	select {}
}

package main

import (
	"fmt"

	_ "github.com/cockroachdb/cockroach-go/v2/crdb"

	"unitasks.josefjantzen.de/backend/api"
	"unitasks.josefjantzen.de/backend/config"
	"unitasks.josefjantzen.de/backend/database"
)

func main() {
	config, err := config.Read("config.sample.json")
	if err != nil {
		fmt.Println("Config read error: ", err)
		return
	}
	dbService := database.InitDB(config)
	api.InitServer(dbService, config)
	fmt.Println("Completly Started")
	select {}
}

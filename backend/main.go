package main

import (
	"errors"
	"fmt"
	"os"

	_ "github.com/cockroachdb/cockroach-go/v2/crdb"

	"unitasks.josefjantzen.de/backend/api"
	"unitasks.josefjantzen.de/backend/config"
	"unitasks.josefjantzen.de/backend/database"
)

func main() {
	var configName string
	if len(os.Args) > 0 {
		configName = os.Args[1]
	}
	if _, err := os.Stat(configName); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Fallback to sample config because ", configName, " doesn't exists")
		configName = "config.sample.json"
	}

	config, err := config.Read(configName)
	if err != nil {
		fmt.Println("Config read error: ", err)
		return
	}
	dbService := database.InitDB(config)
	api.InitServer(dbService, config)
	fmt.Println("Completly Started")
	select {}
}

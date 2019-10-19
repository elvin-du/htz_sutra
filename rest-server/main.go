package main

import (
	"log"
	"mdu/explorer/common/database"
	"mdu/explorer/rest-server/config"
	"mdu/explorer/rest-server/rest/server"
	_ "mdu/explorer/rest-server/rest/controller"
)

func init()  {
	log.SetFlags(log.Lshortfile)
}

func main() {
	database.DefaultDB.Start(config.DefaultConfig.MongoURI)
	server.DefaultServer.Start()
}

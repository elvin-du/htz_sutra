package main

import (
	"mdu/explorer/common/database"
	"mdu/explorer/data-syncer/config"
	"mdu/explorer/data-syncer/service"
	"mdu/explorer/data-syncer/task"
)

func main() {
	database.DefaultDB.Start(config.DefaultConfig.MongoURI)

	newTask := task.NewTask("update_blocks_txs_msgs_validators", service.UpdateBlocks)
	newTask.Start()

	exit := make(chan int)
	<-exit
}

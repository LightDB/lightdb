package main

import (
	"github.com/Im-Stevemmmmm/fluxdb/database"
)

const (
	defaultPort = 1623
)

func main() {
	nodes := make([]*database.ReplicationNode, 0)
	if node := database.VerifyReplicationNodes(nodes); node != nil {
		panic("invalid replication node: " + *node)
	}

	database.Instance = database.NewDB(nodes)

	database.Instance.LoadData()

	initAPI()
	displayStartupMessage()
}

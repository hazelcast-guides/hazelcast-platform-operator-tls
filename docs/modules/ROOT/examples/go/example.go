package main

import (
	"context"
	"fmt"

	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/cluster"
)

func main() {
	var clusterConfig cluster.Config
	clusterConfig.Network.SetAddresses("<EXTERNAL-IP>")
	clusterConfig.Network.SSL.Enabled = true
	clusterConfig.Network.SSL.ServerName = "example"
	clusterConfig.Network.SSL.SetCAPath("example.crt")

	ctx := context.Background()
	config := hazelcast.Config{Cluster: clusterConfig}
	client, err := hazelcast.StartNewClientWithConfig(ctx, config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successful connection!", client)
}

package backend

import (
	"github.com/arcology-network/component-lib/rpc"
	"github.com/smallnest/rpcx/client"
)

var xclient client.XClient
var xclientGateway client.XClient
var xclientScheduler client.XClient

var zookeeperServers []string

func InitParams(servers []string) {
	zookeeperServers = servers
}

func getXclientForStorage() client.XClient {
	if xclient == nil {
		xclient = rpc.InitZookeeperRpcClient("storage", zookeeperServers)
	}
	return xclient
}

func getXclientForGateway() client.XClient {
	if xclientGateway == nil {
		xclientGateway = rpc.InitZookeeperRpcClient("gateway", zookeeperServers)
	}
	return xclientGateway
}

func getXclientForScheduler() client.XClient {
	if xclientScheduler == nil {
		xclientScheduler = rpc.InitZookeeperRpcClient("scheduler", zookeeperServers)
	}
	return xclientScheduler
}

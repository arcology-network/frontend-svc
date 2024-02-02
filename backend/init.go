package backend

import (
	intf "github.com/arcology-network/streamer/interface"
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
		xclient = intf.InitZookeeperRpcClient("storage", zookeeperServers)
	}
	return xclient
}

func getXclientForGateway() client.XClient {
	if xclientGateway == nil {
		xclientGateway = intf.InitZookeeperRpcClient("gateway", zookeeperServers)
	}
	return xclientGateway
}

func getXclientForScheduler() client.XClient {
	if xclientScheduler == nil {
		xclientScheduler = intf.InitZookeeperRpcClient("scheduler", zookeeperServers)
	}
	return xclientScheduler
}

package backend

import (
	"context"
	"fmt"
	"math/big"

	"github.com/arcology-network/common-lib/types"
)

// curl 'http://127.0.0.1:8080/latestheight?access_token=access_token'
func GetLatestHeight() (int, error) {
	request := types.QueryRequest{
		QueryType: types.QueryType_LatestHeight,
	}
	response := types.QueryResult{}

	err := getXclientForStorage().Call(context.Background(), "Query", &request, &response)
	if err != nil {
		return -1, err
	}
	return response.Data.(int), nil
}

// curl 'http://127.0.0.1:8080/nonces/0xcB78F5E0F66bcda91C2899b69Ef703E57C18DCDd?access_token=access_token&height=20'
func GetNonce(address string, height int) (uint64, error) {
	request := types.QueryRequest{
		QueryType: types.QueryType_Nonce,
		Data: types.RequestBalance{
			Height:  height,
			Address: address,
		},
	}
	response := types.QueryResult{}
	err := getXclientForStorage().Call(context.Background(), "Query", &request, &response)
	if err != nil {
		return 0, err
	}
	return response.Data.(uint64), nil
}

// curl 'http://127.0.0.1:8080/balances/0xcB78F5E0F66bcda91C2899b69Ef703E57C18DCDd?access_token=access_token&height=20'
func GetBalance(address string, height int) (*big.Int, error) {
	request := types.QueryRequest{
		QueryType: types.QueryType_Balance,
		Data: types.RequestBalance{
			Height:  height,
			Address: address,
		},
	}
	response := types.QueryResult{}
	err := getXclientForStorage().Call(context.Background(), "Query", &request, &response)
	if err != nil {
		return nil, err
	}
	return response.Data.(*big.Int), nil
}

// curl 'http://127.0.0.1:8080/blocks/944?access_token=access_token&transactions=true'
func GetBlock(height int, transactions bool) (*types.Block, error) {
	request := types.QueryRequest{
		QueryType: types.QueryType_Block,
		Data: &types.RequestBlock{
			Height:       height,
			Transactions: transactions,
		},
	}

	response := types.QueryResult{}
	err := getXclientForStorage().Call(context.Background(), "Query", &request, &response)
	if err != nil {
		return nil, err
	}

	data := response.Data.(types.Block)
	return &data, nil
}

// curl 'http://127.0.0.1:8080/receipts/55312c43a51680df3ec62113c6e0122690b7724cad285e23df8bd01e6f063211?access_token=access_token&executingDebugLogs=true'
// curl 'http://127.0.0.1:8080/receipts/a95c3f8ef0e1fa0855a62f7e00e5121124ca239e2b67b576d0e14672e7507c2e,d3dd60a8dfe4d50c6917fd7bd3ad5226442b25f42110f0e74ecac83adb9913c2?access_token=access_token'
func GetReceipts(hashes []string, executingDebugLogs bool) ([]*types.Receipt, error) {
	request := types.QueryRequest{
		QueryType: types.QueryType_Receipt,
		Data: &types.RequestReceipt{
			Hashes:             hashes,
			ExecutingDebugLogs: executingDebugLogs,
		},
	}
	response := types.QueryResult{}
	err := getXclientForStorage().Call(context.Background(), "Query", &request, &response)
	if err != nil {
		return nil, err
	}
	receipts := response.Data.([]*types.Receipt)
	//receipts := response.Data.(map[string]types.Receipt)
	return receipts, nil
}

// curl 'http://127.0.0.1:8080/containers/0x0000000000000000000000000000000000010203/6601/0?type=array&access_token=access_token&height=20'
// curl 'http://127.0.0.1:8080/containers/0x0000000000000000000000000000000000010203/7801/0?type=queue&access_token=access_token&height=20'
// curl 'http://127.0.0.1:8080/containers/0x0000000000000000000000000000000000010203/7031/00000000000000000000000000000031?type=map&access_token=access_token&height=20'
// curl 'http://127.0.0.1:8080/containers/0x3466323343376365433736314643353663303942/6d617033/0000000000000000000000000000000000000000000000000000000000000026?type=map&access_token=access_token&height=230'
func GetContainer(address, id, key, typ string, height int) (string, error) {
	request := types.QueryRequest{
		QueryType: types.QueryType_Container,
		Data: types.RequestContainer{
			Height:  height,
			Address: address,
			Id:      id,
			Style:   typ,
			Key:     key,
		},
	}
	response := types.QueryResult{}
	err := getXclientForStorage().Call(context.Background(), "Query", &request, &response)
	if err != nil {
		return "", err
	}

	datas := response.Data.([]byte)
	return fmt.Sprintf("%x", datas), nil
}

// curl -i -X POST -H "'Content-type':'application/json'"  'http://127.0.0.1:8080/txs?access_token=access_token&tx=0102030405&tx=0607080910'
func SendTransactions(txs [][]byte) error {
	err := getXclientForGateway().Call(context.Background(), "ReceivedTransactions", &types.SendTransactionArgs{Txs: txs}, &types.SendTransactionReply{})
	if err != nil {
		return err
	}
	return nil
}

// curl -i -X POST -H "'Content-type':'application/json'"  http://127.0.0.1:8080/config?access_token=access_token&parallelism=25
func SetParallelism(p int) (int, error) {
	request := types.ClusterConfig{
		Parallelism: p,
	}
	response := types.SetReply{}

	err := getXclientForScheduler().Call(context.Background(), "SetParallelism", &request, &response)
	if err != nil {
		return 0, err
	}
	return p, nil
}

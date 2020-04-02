package main

import (
	"fmt"
	"testing"
)
const nodeID="3000"
const to="1ekdqSZeFQVsZYJ3CXMMyH6xTjJsU7BDV"
func TestMineReward(t *testing.T) {
	//bc := NewBlockchain("3000")
	balance1 := GetBalance(to, nodeID)
	//fmt.Println(balance1,bc)
	fmt.Sprintf("%d",balance1)

	//var txs []*Transaction
	//
	//cbTx := NewCoinbaseTX( to,"")
	//txs = append(txs, cbTx)
	//
	//newBlock := bc.MineBlock(txs)
	//UTXOSet := UTXOSet{bc}
	//UTXOSet.Reindex()
	//
	//fmt.Println("获取最新出块------newBlock",newBlock)
	////获取账户余额
	//
	//balance2 := GetBalance(to, nodeID)
	//
	//assert.Equal(t,balance1,balance2)

}

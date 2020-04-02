package main

import "fmt"

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)
	//wallet := wallets.GetWallet(address)
	//pubkey->base58->key
	fmt.Printf("Your new address: %s\n",address)
	//fmt.Printf("Your new private-key: %x\n",wallet.PrivateKey)
	//fmt.Printf("Your new public-key: %s\n",wallet.PublicKey)
}



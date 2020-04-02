package main

import (
	"fmt"
	"testing"
)

func TestGetWallet(t *testing.T)  {
	wallets, _ := NewWallets("3000")
	wallet := wallets.GetWallet("1GtuXVRXSkXbsets6wTyRs7R9qWz3KFzu3crypto")
	fmt.Println(wallet)
}

package main

//func TestCreateWallet(t *testing.T) {
//	var cli CLI
//	cli.createWallet("3000")
//}

import (
	"fmt"
	"testing"
)


func TestA(t *testing.T) {
	s := "尹正杰到此一游"
	gbk, err := Utf8ToGbk([]byte(s))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("以GBK的编码方式打开:",string(gbk))
	}

	utf8, err := GbkToUtf8(gbk)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("以UTF8的编码方式打开:",string(utf8))
	}

	var  cli  CLI
	cli.createWallet("3000")

	//wallet := NewWallet()

	//fmt.Println("0 - Having a private ECDSA key")
	//fmt.Println(byteString(wallet.PrivateKey))
	//fmt.Println("=======================")
	//// fmt.Println("private wallet import format")
	//// fmt.Println("private wallet import format", ToWIF(wallet.PrivateKey))
	//// fmt.Println("=======================")
	//fmt.Println("1 - Take the corresponding public key generated with it (65 bytes, 1 byte 0x04, 32 bytes corresponding to X coordinate, 32 bytes corresponding to Y coordinate)")
	//fmt.Println("raw public key", byteString(wallet.PublicKey))
	//fmt.Println("=======================")
	//wallet.GetAddress()
}
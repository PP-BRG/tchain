package cli

import (
	"fmt"
	"log"
	"tchain/wallet"
)

func (cli *CLI) listAddresses() {
	wallets, err := wallet.NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}

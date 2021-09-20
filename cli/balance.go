package cli

import (
	"bitacoin_client/config"
	"bitacoin_client/helper"
	"bitacoin_client/interactor"
	"bitacoin_client/repository"
	"flag"
	"fmt"
	"log"
)

func balance(args []string) error {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	pubKeyFilePath := fs.String("pub", "", "wallet public key file path")
	configFilePath := fs.String("config", "", "path of config file")
	err := fs.Parse(args[1:])
	if err != nil {
		return err
	}

	err = config.ReadConfig(*configFilePath)
	if err != nil {
		log.Fatalf("could not read the config file, err: %s\n")
	}

	if *pubKeyFilePath == "" {
		return fmt.Errorf("pub parameter is required")
	}

	pubKey, err := helper.ReadKeyFromPemFile(*pubKeyFilePath)
	if err != nil {
		return fmt.Errorf("read pub key err: %s", err.Error())
	}

	balanceAmount := 0
	for i := range config.Config.Hosts {
		balanceAmount, err = interactor.GetBalance(pubKey, config.Config.Hosts[i]+repository.BalanceUrl)
		if err != nil {
			log.Printf("get balanceAmount err: %s\n", err.Error())
		} else {
			break
		}
	}

	if balanceAmount != -1 {
		log.Printf("The balanceAmount is %d\n", balanceAmount)
	}

	return nil
}

func init() {
	addCommand("balance", "print balance of your wallet", balance)
}

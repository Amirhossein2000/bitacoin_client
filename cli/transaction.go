package cli

import (
	"bitacoin_client/config"
	"bitacoin_client/helper"
	"bitacoin_client/interactor"
	"flag"
	"fmt"
	"log"
)

func transaction(args []string) error {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	configFilePath := fs.String("config", "", "path of config file")
	PubKeyFilePath := fs.String("pub", "", "source wallet public key file path")
	PrKeyFilePath := fs.String("pr", "", "source wallet private key file path")
	toPubKeyFilePath := fs.String("to", "", "destination wallet public key file path")
	amount := fs.Int("a", 0, "amount of txn")

	err := fs.Parse(args[1:])
	if err != nil {
		return err
	}

	err = config.ReadConfig(*configFilePath)
	if err != nil {
		log.Fatalf("could not read the config file, err: %s\n")
	}

	PrKey, err := helper.ReadKeyFromPemFile(*PrKeyFilePath)
	if err != nil {
		return fmt.Errorf("could not open PrKeyFilePath file, err: %s\n", err.Error())
	}

	pubKey, err := helper.ReadKeyFromPemFile(*PubKeyFilePath)
	if err != nil {
		return fmt.Errorf("could not open PubKeyFilePath file, err: %s\n", err.Error())
	}

	err = helper.VerifyKeys(PrKey, pubKey)
	if err != nil {
		return fmt.Errorf("source private and public key pairs are not valid, err: %s\n", err.Error())
	}

	toPubKey := []byte{}
	if toPubKeyFilePath != nil {
		toPubKey, err = helper.ReadKeyFromPemFile(*toPubKeyFilePath)
		if err != nil {
			return fmt.Errorf("could not open toPubKeyFilePath file, err: %s\n", err.Error())
		}
	} else {
		return fmt.Errorf("there is no destination public key")
	}

	txnReq, err := helper.CreateTxnReq(PrKey, pubKey, toPubKey, *amount)
	if err != nil {
		return fmt.Errorf("could not create new txn, err: %s\n", err.Error())
	}

	interactor.Shout(txnReq, config.Config.Hosts)

	return nil
}

func init() {
	addCommand("transaction", "Print balance for someone", transaction)
}

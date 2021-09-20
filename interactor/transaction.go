package interactor

import (
	"bitacoin_client/helper"
	"bitacoin_client/repository"
	"bitacoin_client/schema"
	"log"
	"net/http"
)

func Shout(txn *schema.TransactionRequest, hosts []string) {
	for i := range hosts {
		err := helper.SendReqAndUnmarshalResp(
			http.MethodPost,
			hosts[i]+repository.TransactionUrl,
			txn,
			http.StatusOK,
			nil,
		)
		if err != nil {
			log.Printf("an error while sending txn to other nodes, node ip: %s err: %s\n", hosts[i], err.Error())
			continue
		}
	}

	log.Println("transaction has been sent to other nodes\n")
}

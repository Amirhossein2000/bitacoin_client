package interactor

import (
	"bitacoin_client/helper"
	"fmt"
	"net/http"
)

func GetBalance(pubKey []byte, url string) (int, error) {
	resp := make(map[string]int)
	err := helper.SendReqAndUnmarshalResp(
		http.MethodPost,
		url,
		map[string][]byte{"pubKey": pubKey},
		http.StatusOK,
		&resp)
	if err != nil {
		return -1, fmt.Errorf("SendRequestThenUnmarshalResponse err: %s", err.Error())
	}

	balance := resp["balance"]
	if balance <= 0 {

	}

	return balance, nil
}

package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendReqAndUnmarshalResp(method, url string, reqBody interface{}, expectedStatusCode int, respStruct interface{}) error {
	var (
		err         error
		jsonReqBody []byte
	)

	if reqBody != nil {
		jsonReqBody, err = json.Marshal(reqBody)
		if err != nil {
			return err
		}
	}

	buf := bytes.NewBuffer(jsonReqBody)
	request, err := http.NewRequest(method, url, buf)
	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != expectedStatusCode {
		return fmt.Errorf("recienve bad status code, recived %d but expcted %d, url: %s, respbody: %s",
			response.StatusCode, expectedStatusCode, url, respBody)
	}

	if respStruct != nil {
		err = json.Unmarshal(respBody, &respStruct)
		if err != nil {
			return fmt.Errorf("unmarshal response err: %s, respBody: %s", err.Error(), reqBody)
		}
	}

	return nil
}

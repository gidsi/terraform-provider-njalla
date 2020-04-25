package njalla

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type ErrorErrorResponse struct {
	Code	int64	`json:"code"`
	Message	string	`json:"message"`
}

type ErrorResponse struct {
	JsonRpc	string	`json:"jsonrpc"`
	Error 	ErrorErrorResponse	`json:"error"`
}

type NjallaClient struct {
	Token	string
	Url		string
}

func (c NjallaClient) DoRequest(method string, params interface{}, response interface{}) error {
	request := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Url, bytes.NewBuffer(requestJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Njalla " + c.Token)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("server returned a non-ok status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result ErrorResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	if result.Error.Code != 0 {
		return errors.New(result.Error.Message)
	}

	if response != nil {
		err = json.Unmarshal(body, response)
		if err != nil {
			return err
		}
	}

	return nil
}

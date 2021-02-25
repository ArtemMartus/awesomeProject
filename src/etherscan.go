package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type etherscanBlockResponse struct {
	Result struct {
		Transactions []struct {
			Value string
		}
	}
}

func parseEtherscanBlockResponse(bodyBytes []byte) (BlockTotalData, error) {
	var parsed etherscanBlockResponse
	err := json.Unmarshal(bodyBytes, &parsed)
	if err != nil {
		println("error translating etherscan response to structure", err.Error())
		println(string(bodyBytes))
		return BlockTotalData{}, err
	}

	transactions := len(parsed.Result.Transactions)
	sum := big.NewFloat(0)
	weiCoef := big.NewFloat(1e-18)

	for _, v := range parsed.Result.Transactions {
		numberStr := strings.Replace(v.Value, "0x", "", -1)
		value, success := new(big.Int).SetString(numberStr, 16)
		if !success {
			return BlockTotalData{}, errors.New("something went wrong with parsing transaction value")
		}
		fValue := new(big.Float).SetInt(value)
		fValue.Mul(fValue, weiCoef)
		sum.Add(sum, fValue)
	}
	f64sum, _ := sum.Float64()
	return BlockTotalData{Transactions: transactions, Amount: f64sum}, err
}

func getBlockFromEtherscan(blockId int) (BlockTotalData, error) {
	if GetConfig().GetToken() == "YourApiKeyToken" {
		time.Sleep(time.Second * 5) // ? default token limit
	} else {
		time.Sleep(time.Second / 5) // api limit
	}

	req, err := http.NewRequest("GET", "https://api.etherscan.io/api", nil)
	if err != nil {
		return BlockTotalData{}, err
	}

	q := req.URL.Query()
	q.Add("module", "proxy")
	q.Add("action", "eth_getBlockByNumber")
	q.Add("tag", fmt.Sprintf("0x%x", blockId))
	q.Add("boolean", "true")
	q.Add("apikey", GetConfig().GetToken())

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	resp, err := GetConfig().GetHttpClient().Do(req)
	if err != nil {
		println("error happened in etherscan get request", err.Error())
		return BlockTotalData{}, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("error reading response from etherscan", err.Error())
		return BlockTotalData{}, err
	}

	return parseEtherscanBlockResponse(bodyBytes)
}

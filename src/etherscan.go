package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
		return BlockTotalData{}, err
	}

	transactions := len(parsed.Result.Transactions)
	sum := 0.0

	for _, v := range parsed.Result.Transactions {
		numberStr := strings.Replace(v.Value, "0x", "", -1)
		value, _ := strconv.ParseInt(numberStr, 16, 64)
		sum += float64(value) * 1e-18
	}
	return BlockTotalData{Transactions: transactions, Amount: sum}, err
}

func getBlockFromEtherscan(blockId int) (BlockTotalData, error) {
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
		println("error happened in etherscan get request", err)
		return BlockTotalData{}, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return BlockTotalData{}, err
	}

	return parseEtherscanBlockResponse(bodyBytes)
}

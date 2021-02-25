package tests

import (
	"awesomeProject/src"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func blockHelper(blockId int, srv *httptest.Server) (src.BlockTotalData, error) {
	res, err := http.Get(fmt.Sprintf("%s/api/block/%d/total", srv.URL, blockId))

	if err != nil {
		return src.BlockTotalData{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return src.BlockTotalData{}, err
	}

	var data src.BlockTotalData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return src.BlockTotalData{}, err
	}

	return data, err
}

func TestBlockTotal(t *testing.T) {
	srv := httptest.NewServer(src.Handlers())
	defer srv.Close()

	data, err := blockHelper(11509797, srv)
	if err != nil {
		t.Error(err)
	}
	if data.Transactions != 155 {
		t.Fail()
	}
	if data.Amount < 2.285404 || data.Amount > 2.285406 {
		t.Fail()
	}

	data, err = blockHelper(11508993, srv)
	if err != nil {
		t.Error(err)
	}
	if data.Transactions != 241 {
		t.Fail()
	}
	if data.Amount < 1130.987084 || data.Amount > 1130.987086 {
		t.Fail()
	}

	data, err = blockHelper(109789, srv)
	if err != nil {
		t.Error(err)
	}
	if data.Transactions != 1 {
		t.Fail()
	}
	if data.Amount < 4.99876 || data.Amount > 4.99878 {
		t.Fail()
	}
}

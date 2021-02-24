package tests

import (
	"awesomeProject/src"
	"encoding/json"
	"strings"
	"testing"
)

func TestBlockTotalDataTransactions(t *testing.T) {
	data := src.BlockTotalData{Transactions: 100, Amount: 1.123}
	json_ed, _ := json.Marshal(data)

	strJson := string(json_ed)
	if !strings.Contains(strJson, "\"transactions\"") {
		t.Fail()
	}
	if !strings.Contains(strJson, "100") {
		t.Fail()
	}
}

func TestBlockTotalDataAmount(t *testing.T) {
	data := src.BlockTotalData{Transactions: 100, Amount: 1.123}
	json_ed, _ := json.Marshal(data)

	strJson := string(json_ed)

	if !strings.Contains(strJson, "\"amount\"") {
		t.Fail()
	}
	if !strings.Contains(strJson, "1.123") {
		t.Fail()
	}
}

package tests

import (
	"awesomeProject/src"
	"testing"
)

func TestCacheItemExist(t *testing.T) {
	c := src.GetConfig().GetCache()
	if c.CheckBlockData(1) != nil {
		t.Fail()
	}
	if c.CheckBlockData(0) != nil {
		t.Fail()
	}
}

func TestCacheItemInsert(t *testing.T) {
	c := src.GetConfig().GetCache()
	c.InsertBlockData(0, &src.BlockTotalData{
		Transactions: 1,
		Amount:       2,
	})
	c.InsertBlockData(1, &src.BlockTotalData{
		Transactions: 2,
		Amount:       3,
	})

	a := c.CheckBlockData(0)
	aa := c.CheckBlockData(0)
	if a != aa {
		t.Fail()
	}
	if a == nil || a.Transactions != 1 {
		t.Fail()
	}
	b := c.CheckBlockData(1)
	if b == nil || b.Transactions != 2 {
		t.Fail()
	}

}

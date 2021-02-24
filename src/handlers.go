package src

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func BlockTotalHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	blockId, err := strconv.Atoi(ps.ByName("blockId"))
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprintf(w, "{\"error\"=\"bad block id\"}")
		println(err)
		return
	}
	cached := GetConfig().GetCache().CheckBlockData(blockId)
	if cached == nil {
		println("downloading data")
		data, err := getBlockFromEtherscan(blockId)
		if err != nil {
			fmt.Fprint(w, "{\"error\"=\"Something went wrong with etherscan\"}")
			return
		}
		cached = &data
		GetConfig().GetCache().InsertBlockData(blockId, cached)
	} else {
		println("using cached data")
	}

	jData, err := json.Marshal(*cached)
	if err != nil {
		fmt.Fprintf(w, "{\"error\"=\"Something went wrong \\( o_o) / /.. .. .\")")
		println(err)
		return
	}

	w.Write(jData)
}

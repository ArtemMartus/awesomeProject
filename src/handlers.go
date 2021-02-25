package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func Handlers() http.Handler {
	out := http.NewServeMux()

	out.HandleFunc("/", IndexHandler)
	out.HandleFunc("/api/block/", BlockTotalHandler)
	return out
}

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "Welcome!\n")
}

func BlockTotalHandler(w http.ResponseWriter, r *http.Request) {
	validBlockId := regexp.MustCompile("/api/block/(?P<blockId>\\d*)/total")

	strBlockId := validBlockId.FindStringSubmatch(r.RequestURI)
	if len(strBlockId) <= 1 || strBlockId[1] == "" {
		_, _ = fmt.Fprintf(w, "bad url = %s", r.RequestURI)
		println("bad url", r.RequestURI)
		return
	}
	println("blockid = ", strBlockId[1])
	blockId, err := strconv.Atoi(strBlockId[1])
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		_, _ = fmt.Fprintf(w, "{\"error\"=\"bad block id\"}")
		println(err.Error())
		return
	}
	cached := GetConfig().GetCache().CheckBlockData(blockId)
	if cached == nil {
		println("downloading data")
		data, err := getBlockFromEtherscan(blockId)
		if err != nil {
			println("error with etherscan", err.Error())
			_, _ = fmt.Fprint(w, "{\"error\"=\"Something went wrong with etherscan\"}")
			return
		}
		cached = &data
		GetConfig().GetCache().InsertBlockData(blockId, cached)
	} else {
		println("using cached data")
	}

	jData, err := json.Marshal(*cached)
	if err != nil {
		_, _ = fmt.Fprintf(w, "{\"error\"=\"Something went wrong \\( o_o) / /.. .. .\")")
		println(err.Error())
		return
	}

	_, _ = w.Write(jData)
}

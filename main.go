// @Title 
// @Description  
// @Author  Niels  2020/11/20 
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/niels1286/nuls-go-sdk/account"
	"github.com/niels1286/nuls-go-sdk/tx/txdata"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
)

type ResultVo struct {
	Sender          string
	ContractAddress string
	Value           string
	GasLimit        uint64
	Price           uint64
	MethodName      string
	MethodDesc      string
	ArgsCount       uint8
	Args            [][]string
}

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("Parameter wrong!")
	//	return
	//}
	//txDataHex := os.Args[1]
	txDataHex := "0200013452fcc77369361d225f61ea34b9930f98d2b536020002a6d26cb6b330c7a0aaf47848e5c27e3abe1a980400000000000000000000000000000000000000000000000000000000000000007a500000000000001900000000000000087365744167656e74000100"

	txDataBytes, err := hex.DecodeString(txDataHex)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := &txdata.CallContract{}
	data.Parse(seria.NewByteBufReader(txDataBytes, 0))
	result := &ResultVo{
		Sender:          account.GetStringAddress(data.Sender, "NULS"),
		ContractAddress: account.GetStringAddress(data.ContractAddress, "NULS"),
		Value:           data.Value.String(),
		GasLimit:        data.GasLimit,
		Price:           data.Price,
		MethodName:      data.MethodName,
		MethodDesc:      data.MethodDesc,
		ArgsCount:       data.ArgsCount,
		Args:            data.Args,
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(bytes))
}

// @Title
// @Description
// @Author  Niels  2020/11/20
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/niels1286/nuls-go-sdk/account"
	"github.com/niels1286/nuls-go-sdk/crypto/base58"
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
	txDataHex := "0100011d52afe277b0c575355e618a194ffa1ae1cb518f010002be36277487d7c45974d1fcc722d6729f47cddd8900000000000000000000000000000000000000000000000000000000000000001b540000000000001900000000000000087472616e73666572000201254e554c5364364867554763523351626f76593155467561753532374532464e354c70434e54010a33303030303030303030"

	txDataBytes, err := hex.DecodeString(txDataHex)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := &txdata.CallContract{}
	data.Parse(seria.NewByteBufReader(txDataBytes, 0))
	result := &ResultVo{
		Sender:          GetStringAddress(data.Sender, "NULS"),
		ContractAddress: GetStringAddress(data.ContractAddress, "NULS"),
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

//根据地址字节数组，生成可以阅读的字符串地址
func GetStringAddress(bytes []byte, prefix string) string {
	//将之前得到的所有字节，进行异或操作，得到结果追加到
	xor := calcXor(bytes)
	newbytes := []byte{}
	newbytes = append(newbytes, bytes...)
	newbytes = append(newbytes, xor)
	return prefix + account.PrefixTable[len(prefix)] + base58.Encode(newbytes)
}

//计算异或字节
func calcXor(bytes []byte) byte {
	xor := byte(0)
	for _, one := range bytes {
		xor ^= one
	}
	return xor
}

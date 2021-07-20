package models

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/ybbus/jsonrpc"
)

func callContract(ContractAddr, archNode string) (string, error)  {
	//init a new json rpc client
	client := jsonrpc.NewClient(archNode)

	//to assemble the data string structure with fn prefix, addr with left padding
	validatorContractAddr := ContractAddr

	//todo integrate with contract, fnSig and args
	//fn distribution signature in smart contract
	fnSig := "0x000000"
	args := "000000000000000000000000"
	dataOb := fnSig + args

	resp, err := client.Call("eth_call",map[string]interface{}{
		"to": validatorContractAddr,
		"data": dataOb,
	})
	if err != nil {
		logrus.Errorf("eth call error with %v", err)
		return "",errors.New("eth call error")
	}

	result, err := resp.GetString()
	if err != nil {
		return "",errors.New("eth call error")
	}
	return result, nil
}

//removeConZero
func removeConZero(str string) (string) {
	var index int
	sb := []byte(str)
	for i := 0; i < len(sb); i++ {
		if sb[i] == 48 {
			continue
		} else {
			index = i
			break
		}
	}
	return str[index:]
}
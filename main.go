package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://api.shasta.trongrid.io/wallet/deploycontract"

	payload := strings.NewReader("{\"owner_address\":\"41e7ebcff07619ed7956b79b9a8c13d31deb5f8ac1\",\"abi\":\"[{\\\"inputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"ch\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"dbc\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"constructor\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"counta\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"countb\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"c\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"testA\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"pure\\\",\\\"type\\\":\\\"function\\\"}]\",\"bytecode\":\"608060405234801561001057600080fd5b50d3801561001d57600080fd5b50d2801561002a57600080fd5b506004361061005b5760003560e01c806348e2a46e14610060578063782d41221461007e5780639d4853c41461009c575b600080fd5b6100686100cc565b6040516100759190610133565b60405180910390f35b6100866100d2565b6040516100939190610133565b60405180910390f35b6100b660048036038101906100b191906100f7565b6100d8565b6040516100c39190610133565b60405180910390f35b60005481565b60015481565b6000819050919050565b6000813590506100f18161015d565b92915050565b60006020828403121561010d5761010c610158565b5b600061011b848285016100e2565b91505092915050565b61012d8161014e565b82525050565b60006020820190506101486000830184610124565b92915050565b6000819050919050565b600080fd5b6101668161014e565b811461017157600080fd5b5056fea26474726f6e58221220373310d52a7dbaffda61e3931b76001f11ead0d60def6ca8e0b8d7bad5fbdc4964736f6c63430008060033\",\"fee_limit\":1000000,\"origin_energy_limit\":100000,\"name\":\"AaaContract\",\"call_value\":0,\"consume_user_resource_percent\":100,\"parameter\":\"00000000000000000000000000000000000000000000000000000000000000420000000000000000000000000000000000000000000000000000000000000002\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}

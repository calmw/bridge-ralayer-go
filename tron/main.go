package main

import (
	"bridge-relayer/tron/lib"
	"bridge-relayer/utils"
	"fmt"
)

var bridgeAddress = "TSUbUczwFjPt6vweMT4rrJ7yL1mzjaaEqk"
var OwnerAddress = "TX7VatsLxHP9VhwxjtraQmbsDaBNf8ykpW"

func main() {
	address := "ca98cd464a9cc69409de5bfec896a45985f670a3"
	base58, err := utils.HexToBase58(&address)
	fmt.Println(*base58, err)
	addressHex, err := utils.Base58ToHex(base58)
	fmt.Println(*addressHex, err)
	target := "ca98cd464a9cc69409de5bfec896a45985f670a3"
	base58Target, err := utils.HexToBase58(&target)
	fmt.Println(*base58Target, err)

	lib.CallRemote("3565375400000000000000000000000000000000000000000000000000000000", "1", *base58, "3565375400000000000000000000000000000000000000000000000000000000")
	//
	//resourceIDBytes, _ := hex.DecodeString("3565375400000000000000000000000000000000000000000000000000000000")
	//resourceID := utils.ByteSliceToByte32(resourceIDBytes)
	//
	//messageIDBytes, _ := hex.DecodeString("155c490c849f5126b70449ab9366c098c681f7df3900d8e432ae5bcb59d424ad")
	//messageID := utils.ByteSliceToByte32(messageIDBytes)
	//
	////dataHashBytes, _ := hex.DecodeString("155c490c849f5126b70449ab9366c098c681f7df3900d8e432ae5bcb59d424ad")
	////dataHash := utils.ByteSliceToByte32(dataHashBytes)
	//
	//data, _ := hex.DecodeString("3565375400000000000000000000000000000000000000000000000000000000")
	//
	//sig1, _ := hex.DecodeString("66dbff2b9394fe145cdaa8a1fa284801f785e5b6ca2e7f35cd563572d6dca5fc2fb6812e036bd6a6c1bb59c895ca4525ffefe2b60448787f8c607b0d0ecde5961c")
	//sig2, _ := hex.DecodeString("309b66297594ede67ff764cf98e892da5ddc20d39c3c6ab15afd908f3a7da0231102cf5d2d000cc484fc092ece37d49f5a4e3b67f764510eead316a6bc96601f1c")
	//var sig [][]byte
	////fmt.Println(sig2)
	//sig = append(sig, sig1)
	//sig = append(sig, sig2)
	//pack, err := lib.Pack("execute(bytes32,uint32,uint256,bytes32,uint32,address,bytes,bytes[])",resourceID,messageID, dataHash,data)
	//fmt.Println(err,pack)
	//if err != nil {
	//	return
	//}

	//request := trigger.ExecuteRequest{
	//	ResourceID:    resourceID, //  [32]byte
	//	SourceChainId: 6,
	//	SourceNonce:   big.NewInt(6), //*big.Int
	//	TargetChainId: 1,             // uint32
	//	Target:        common.HexToAddress("ca98cd464a9cc69409de5bfec896a45985f670a3"),
	//	MessageId:     messageID, // [32]byte
	//	DataHash:      dataHash,  //   [32]byte
	//	Data:          data,      //[]byte
	//	Signatures:    sig,       //[][]byte
	//}

	//lib.Execute(trigger.ExecuteRequest{
	//	ResourceID:    resourceID, //  [32]byte
	//	SourceChainId: 6,
	//	SourceNonce:   big.NewInt(6), //*big.Int
	//	TargetChainId: 1,             // uint32
	//	Target:        address3.Address{},
	//	MessageId:     messageID, // [32]byte
	//	DataHash:      dataHash,  //   [32]byte
	//	Data:          data,      //[]byte
	//	Signatures:    sig,       //[][]byte
	//})
	//ethCli, err := ethclient.Dial("https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	//fmt.Println(err)
	//newBridge, err := bridge.NewBridge(common.HexToAddress("0x451e2ede33b13ff00a3d28f8565032553bf2b486"), ethCli)
	//fmt.Println(err)
	//if err != nil {
	//	return
	//}
	//contractAbi, err := bridge.BridgeMetaData.GetAbi()
	//pack, err := contractAbi.Pack("execute", resourceID,uint32(6),big.NewInt(6),messageID,uint32(1),common.HexToAddress("ca98cd464a9cc69409de5bfec896a45985f670a3"),data,sig)
	//fmt.Println(err,pack)

	//b, err := hex.DecodeString("ca98cd464a9cc69409de5bfec896a45985f670a3")
	//
	//hash0 := utils.Hash(b)
	//hash1 := utils.Hash(hash0)
	//// Since hash (sha256) never fails, the hash should always have length >4
	//inputCheck := append(b, hash1[:4]...)
	//fmt.Println(inputCheck)
	//
	//contractAbi, err := bridge.BridgeMetaData.GetAbi()
	//pack, err := contractAbi.Pack("execute", resourceID,uint32(6),big.NewInt(6),messageID,uint32(1),utils.ByteSliceToByte32(inputCheck),data,sig)
	//fmt.Println(err,pack)
	//if err != nil {
	//	return
	//}

	//lib.Execute("3565375400000000000000000000000000000000000000000000000000000000", "6", "6",
	//	"155c490c849f5126b70449ab9366c098c681f7df3900d8e432ae5bcb59d424ad", "1",
	//	*base58Target, "3565375400000000000000000000000000000000000000000000000000000000", []string{
	//		"0x66dbff2b9394fe145cdaa8a1fa284801f785e5b6ca2e7f35cd563572d6dca5fc2fb6812e036bd6a6c1bb59c895ca4525ffefe2b60448787f8c607b0d0ecde5961c",
	//		"0x309b66297594ede67ff764cf98e892da5ddc20d39c3c6ab15afd908f3a7da0231102cf5d2d000cc484fc092ece37d49f5a4e3b67f764510eead316a6bc96601f1c",
	//	})

	//jsonData := `[{"bytes32":"3565375400000000000000000000000000000000000000000000000000000000"},{"uint32":"6"},{"uint256":"6"},{"uint32":"1"},{"address":"TX7VatsLxHP9VhwxjtraQmbsDaBNf8ykpW"},{"bytes32":"3565375400000000000000000000000000000000000000000000000000000000"},{"bytes32":"3565375400000000000000000000000000000000000000000000000000000000"},{"bytes[]":"[66dbff2b9394fe145cdaa8a1fa284801f785e5b6ca2e7f35cd563572d6dca5fc2fb6812e036bd6a6c1bb59c895ca4525ffefe2b60448787f8c607b0d0ecde5961c,309b66297594ede67ff764cf98e892da5ddc20d39c3c6ab15afd908f3a7da0231102cf5d2d000cc484fc092ece37d49f5a4e3b67f764510eead316a6bc96601f1c]"}]`
	//err = lib.Execute(jsonData)
	//fmt.Println(err, 9999)

	//err = lib.Execute2(pack)
	//fmt.Println(err, 9999)
}

//genericKeyStore.GenericKeysFromPrivateKey("81290e5eac98cb77f265ff0cd92b11619ba59d95404c14dc1c66b4e42c0aca18")
// err := lib.Deploy("bridge")
//lib.CallRemote("3565375400000000000000000000000000000000000000000000000000000000","1","0xca98cd464a9cc69409de5bfec896a45985f670a3","3565375400000000000000000000000000000000000000000000000000000000")
//lib.Execute("cccc", "cccc", "cccc", "cccc", "cccc", "cccc", "cccc", []string{"aaaaa", "bbbbb"})
//fmt.Println(lib.GetLatestBlock())

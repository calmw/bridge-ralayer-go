package tron

import (
	"bridge-relayer/utils"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var DeployUri string = "https://api.shasta.trongrid.io/wallet/deploycontract"
var Web3Uri string = "https://api.shasta.trongrid.io/v1/contracts/"
var GetNowBlockUri string = "https://api.shasta.trongrid.io/walletsolidity/getnowblock"

type Deploy struct {
	OwnerAddress               string `json:"owner_address"`                 // 合约部署者地址，Base58check 格式或 HEX 格式
	Abi                        string `json:"abi"`                           // 智能合约的 ABI
	Bytecode                   string `json:"bytecode"`                      // 合约二进制代码, HEX 格式
	FeeLimit                   int    `json:"fee_limit"`                     // 部署合约最大 TRX 消耗量，单位为 SUN（1 TRX = 1,000,000 SUN）
	Parameter                  string `json:"parameter"`                     // 传递给合约构造函数的参数。 调用参数[1、2]的虚拟机格式，使用remix提供的js工具，将合同调用方调用的参数数组[1、2]转换为虚拟机所需的参数格式。
	OriginEnergyLimit          int    `json:"origin_energy_limit"`           // 在执行或创建合同的过程中，所有者将消耗的最大能量是一个应该大于 0 的整数。
	Name                       string `json:"name"`                          // 合约名称, 例如 SomeContract
	CallValue                  int    `json:"call_value"`                    // 部署时候往合约转账的 TRX 数量, 以 SUN 为单位
	ConsumeUserResourcePercent int    `json:"consume_user_resource_percent"` // 指定的使用该合约用户的资源占比，是[0, 100]之间的整数。如果是0，则表示用户不会消耗资源。如果开发者资源消耗完了，才会完全使用用户的资源。
}

type DeployResponse struct {
	Visible         bool   `json:"visible"`
	TxID            string `json:"txID"`
	ContractAddress string `json:"contract_address"`
	RawData         struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					OwnerAddress string `json:"owner_address"`
					NewContract  struct {
						Bytecode                   string `json:"bytecode"`
						ConsumeUserResourcePercent int    `json:"consume_user_resource_percent"`
						Name                       string `json:"name"`
						OriginAddress              string `json:"origin_address"`
						Abi                        struct {
							Entrys []struct {
								Inputs []struct {
									Name string `json:"name"`
									Type string `json:"type"`
								} `json:"inputs"`
								Name            string `json:"name"`
								StateMutability string `json:"stateMutability"`
								Type            string `json:"type"`
								Outputs         []struct {
									Name string `json:"name"`
									Type string `json:"type"`
								} `json:"outputs,omitempty"`
								Constant bool `json:"constant,omitempty"`
							} `json:"entrys"`
						} `json:"abi"`
						OriginEnergyLimit int `json:"origin_energy_limit"`
					} `json:"new_contract"`
				} `json:"value"`
				TypeURL string `json:"type_url"`
			} `json:"parameter"`
			Type string `json:"type"`
		} `json:"contract"`
		RefBlockBytes string `json:"ref_block_bytes"`
		RefBlockHash  string `json:"ref_block_hash"`
		Expiration    int64  `json:"expiration"`
		FeeLimit      int    `json:"fee_limit"`
		Timestamp     int64  `json:"timestamp"`
	} `json:"raw_data"`
	RawDataHex string `json:"raw_data_hex"`
}

type EventResponse struct {
	Data    []interface{} `json:"data"`
	Success bool          `json:"success"`
	Meta    struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}

type EventErr struct {
	Success    bool   `json:"success"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

type EventRequest struct {
	EventName         string // 事件名称, 例如 Transfer
	BlockNumber       int
	OnlyUnconfirmed   bool
	OnlyConfirmed     bool
	MinBlockTimestamp int    // 起始区块时间戳, 默认为 0
	OrderBy           string // block_timestamp,desc | block_timestamp,asc
	Limit             int    // 每页结果数, 默认 20 最大 200，如果不需要分页，建议设最大值
	fingerprint       string // 翻页参数, 指定上一页的 fingerprint
}

type NowBlockResponse struct {
	BlockID     string `json:"blockID"`
	BlockHeader struct {
		RawData struct {
			Number         int    `json:"number"`
			TxTrieRoot     string `json:"txTrieRoot"`
			WitnessAddress string `json:"witness_address"`
			ParentHash     string `json:"parentHash"`
			Version        int    `json:"version"`
			Timestamp      int64  `json:"timestamp"`
		} `json:"raw_data"`
		WitnessSignature string `json:"witness_signature"`
	} `json:"block_header"`
	Transactions []struct {
		Ret []struct {
			ContractRet string `json:"contractRet"`
		} `json:"ret"`
		Signature []string `json:"signature"`
		TxID      string   `json:"txID"`
		RawData   struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Data            string `json:"data"`
						OwnerAddress    string `json:"owner_address"`
						ContractAddress string `json:"contract_address"`
						CallValue       int    `json:"call_value"`
					} `json:"value"`
					TypeURL string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			FeeLimit      int    `json:"fee_limit"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
	} `json:"transactions"`
}

type CallContractRequest struct {
	OwnerAddress     string `json:"owner_address"`     // 发起合约调用的账户地址, Base58check 格式或 HEX 格式,41D1E7A6BC354106CB410E65FF8B181C600FF14292
	ContractAddress  string `json:"contract_address"`  // 合约地址, Base58check 格式或 HEX 格式,41a7837ce56da0cbb28f30bcd5bff01d4fe7e4c6e3
	FunctionSelector string `json:"function_selector"` // 所调用的函数, transfer(address,uint256)
	Parameter        string `json:"parameter"`         // parameter的编码需要根据合约的ABI规则进行，规则比较复杂， 用户可以利用ethers库进行编码，具体可以参考指南中智能合约篇参数编码和解码文档。
	FeeLimit         int    `json:"fee_limit"`         // 最大消耗的 TRX 数量, 以 SUN 为单位
	CallValue        int    `json:"call_value"`        // 本次调用往合约转账的 TRX 数量, 以 SUN 为单位
	PermissionId     int    `json:"permission_id"`     // 可选参数，多重签名时使用
	Visible          bool   `json:"visible"`           // 账户地址是否为 Base58check 格式, 默认为 false, 使用 HEX 地址
}

type CallContractResponseErr struct {
	Result struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"result"`
}

type CallConstContractRequest struct {
	OwnerAddress     string `json:"owner_address"`     // 发起合约调用的账户地址, Base58check 格式或 HEX 格式,41D1E7A6BC354106CB410E65FF8B181C600FF14292
	ContractAddress  string `json:"contract_address"`  // 合约地址, Base58check 格式或 HEX 格式,41a7837ce56da0cbb28f30bcd5bff01d4fe7e4c6e3
	FunctionSelector string `json:"function_selector"` // 所调用的函数, transfer(address,uint256)
	Parameter        string `json:"parameter"`         // parameter的编码需要根据合约的ABI规则进行，规则比较复杂， 用户可以利用ethers库进行编码，具体可以参考指南中智能合约篇参数编码和解码文档。
	Visible          bool   `json:"visible"`           // 账户地址是否为 Base58check 格式, 默认为 false, 使用 HEX 地址
}

func DeployContract(deploy Deploy) (error, DeployResponse) {

	deployBytes, _ := json.Marshal(deploy)
	res, err := utils.HttpPost(DeployUri, deployBytes)
	if err != nil {
		return err, DeployResponse{}
	}

	if strings.Contains(string(res), "Error") {
		fmt.Println(errors.New("deploy error"))
	}

	var deployResponse DeployResponse
	err = json.Unmarshal(res, &deployResponse)
	if err != nil {
		return err, DeployResponse{}
	}
	return nil, deployResponse
}

/*
GetEventByBlock
contractAddress TX7VatsLxHP9VhwxjtraQmbsDaBNf8ykpW
*/
func GetEventByBlock(contractAddress string, data EventRequest) (error, EventResponse) {

	header := map[string]string{
		"Accept": "application/json",
	}
	Web3Uri = Web3Uri + contractAddress + "/events"
	dataMap := utils.StructToStringMap(data)
	res, err := utils.HttpGet(Web3Uri, header, dataMap)
	if err != nil {
		return err, EventResponse{}
	}

	if strings.Contains(string(res), "Error") {
		var eventErr EventErr
		err := json.Unmarshal(res, &eventErr)
		if err != nil {
			return err, EventResponse{}
		}
		return errors.New(eventErr.Error), EventResponse{}
	}
	var eventResponse EventResponse
	err = json.Unmarshal(res, &eventResponse)
	if err != nil {
		return err, EventResponse{}
	}

	return nil, eventResponse
}

// GetNowBlock get the latest block number
func GetNowBlock() (error, NowBlockResponse) {

	NowBlock, err := utils.HttpGet(GetNowBlockUri, nil, nil)
	if err != nil {
		return err, NowBlockResponse{}
	}

	var nowBlockResponse NowBlockResponse

	err = json.Unmarshal(NowBlock, &nowBlockResponse)
	if err != nil {
		return err, NowBlockResponse{}
	}

	return nil, nowBlockResponse
}

// CallContract 调用智能合约
func CallContract(data CallContractRequest) (error, NowBlockResponse) {

	NowBlock, err := utils.HttpGet(GetNowBlockUri, nil, nil)
	if err != nil {
		return err, NowBlockResponse{}
	}

	var nowBlockResponse NowBlockResponse

	err = json.Unmarshal(NowBlock, &nowBlockResponse)
	if err != nil {
		return err, NowBlockResponse{}
	}

	return nil, nowBlockResponse
}

// CallConstContract 调用合约的常量函数, 需要函数类型为 view 或 pure.
func CallConstContract(data CallConstContractRequest) (error, NowBlockResponse) {

	NowBlock, err := utils.HttpGet(GetNowBlockUri, nil, nil)
	if err != nil {
		return err, NowBlockResponse{}
	}

	var nowBlockResponse NowBlockResponse

	err = json.Unmarshal(NowBlock, &nowBlockResponse)
	if err != nil {
		return err, NowBlockResponse{}
	}

	return nil, nowBlockResponse
}

package webThree

var Web3Uri string = "https://api.shasta.trongrid.io/wallet/triggersmartcontract"

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

type CallContractResponse struct {
	Result struct {
		Result bool `json:"result"`
	} `json:"result"`
	EnergyUsed     int      `json:"energy_used"`
	ConstantResult []string `json:"constant_result"`
	Transaction    struct {
		Ret []struct {
		} `json:"ret"`
		Visible bool   `json:"visible"`
		TxID    string `json:"txID"`
		RawData struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Data            string `json:"data"`
						OwnerAddress    string `json:"owner_address"`
						ContractAddress string `json:"contract_address"`
					} `json:"value"`
					TypeUrl string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
	} `json:"transaction"`
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

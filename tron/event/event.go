package event

type Block struct {
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
}

type EventResponse struct {
	Data    []EventData `json:"data"`
	Success bool        `json:"success"`
	Meta    struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}

type EventData struct {
	BlockNumber           int         `json:"block_number"`
	BlockTimestamp        int64       `json:"block_timestamp"`
	CallerContractAddress string      `json:"caller_contract_address"`
	ContractAddress       string      `json:"contract_address"`
	EventIndex            int         `json:"event_index"`
	EventName             string      `json:"event_name"`
	Result                interface{} `json:"result"`
	ResultType            interface{} `json:"result_type"`
	Event                 string      `json:"event"`
	TransactionId         string      `json:"transaction_id"`
}

type CallRequestEventLog struct {
	ResourceID    string `json:"resourceID"`
	Data          string `json:"data"`
	MessageId     string `json:"messageId"`
	SourceChainId string `json:"sourceChainId"`
	Target        string `json:"target"`
	Field6        string `json:"0"`
	Field7        string `json:"1"`
	Caller        string `json:"caller"`
	Field9        string `json:"2"`
	Field10       string `json:"3"`
	Field11       string `json:"4"`
	Field12       string `json:"5"`
	SourceNonce   string `json:"sourceNonce"`
	Field14       string `json:"6"`
	TargetChainId string `json:"targetChainId"`
	Field16       string `json:"7"`
}

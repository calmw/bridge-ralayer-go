package event

import (
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Sig string

const (
	CallRequest        Sig = "CallRequest(bytes32,address,uint32,uint256,bytes32,uint32,address,bytes)"
	ConfirmedRequest   Sig = "ConfirmedRequest(bytes32,address,uint32,uint256,bytes32,uint32,address,bytes)"
	SignatureCollected Sig = "SignatureCollected(bytes32,uint8,uint32,uint256,uint32,address,bytes32,bytes32,bytes[])"
)

type Event struct {
	EventSignature Sig
	EventName      string
}

var CallRequestEvent = Event{
	CallRequest,
	"CallRequest",
}
var ConfirmedRequestEvent = Event{
	ConfirmedRequest,
	"ConfirmedRequest",
}
var SignatureCollectedEvent = Event{
	SignatureCollected,
	"SignatureCollected",
}

func (es Sig) GetTopic() ethCommon.Hash {
	return crypto.Keccak256Hash([]byte(es))
}

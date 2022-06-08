package event

import (
	"bridge-relayer/services/event"
	eth "github.com/ethereum/go-ethereum"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"math/big"
)

func BuildQuery(contract ethCommon.Address, sig event.Sig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []ethCommon.Address{contract},
		Topics: [][]ethCommon.Hash{
			{sig.GetTopic()},
		},
	}

	return query
}

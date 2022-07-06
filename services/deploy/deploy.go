package deploy

import (
	"bridge-relayer/binding/bridge"
	"bridge-relayer/internal/relayer"
	"bridge-relayer/services"
	"bridge-relayer/utils"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ContractName = map[string]bool{
	"bridge":  true,
	"manager": true,
}

func DeployContract(contractName, chainId string) {
	log := log15.Root().New("deploy", contractName)
	if contractName == "bridge" {
		ok, chain := services.CheckChainId(utils.StringToInt(chainId))
		if !ok {
			log.Error("deploy contract err", "chain id error")
			return
		}

		ethCli, err := ethclient.Dial(chain.Endpoint)
		if nil != err {
			log.Error("deploy contract err", err.Error())
			return
		}

		contractAbi, err := bridge.BridgeMetaData.GetAbi()
		if nil != err {
			log.Error("deploy contract err", err.Error())
			return
		}
		contract, t, _, err := bind.DeployContract(relayer.ThisReLayer.TransactOpts, *contractAbi, []byte(bridge.BridgeMetaData.Bin), ethCli, uint8(0), uint32(0))
		if err != nil {
			log.Error("deploy contract err", err.Error())
			return
		}
		log.Info("deploy contract success", "contract address", contract, "txHash", t.Hash())
	}
}

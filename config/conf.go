package config

import (
	"bridge-relayer/keyStore"
	"bridge-relayer/utils"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"os"
)

const DefaultGasLimit = 6721975
const DefaultBlockConfirmations = 1

var Config *Conf

var ChainCfg map[int]Chain
var EngineCfg Engine
var ReLayerNum int

type Conf struct {
	Engine EngineConfig
	Chains []ChainsConfig
}

type Engine struct {
	Name           string
	ChainId        int64
	Endpoint       string
	StartBlock     *big.Int
	ManagerAddress common.Address
}

type EngineConfig struct {
	Name           string `toml:"chain_name"`
	ChainId        int64  `toml:"chain_id"`
	Endpoint       string `toml:"endpoint"`
	ManagerAddress string `toml:"manager_address"`
}

type ChainsConfig struct {
	Name       string   `toml:"chain_name"`
	Id         int      `toml:"id"`
	ChainId    int64    `toml:"chain_id"`
	VoteChain  bool     `toml:"vote_chain"`
	Endpoint   string   `toml:"endpoint"`
	Bridge     string   `toml:"bridge_address"`
	PrivateKey string   `toml:"reLayer_private_key"`
	StartBlock *big.Int `toml:"start_block"`
}

type Chain struct {
	Name       string
	Id         int
	ChainId    int64
	VoteChain  bool
	Endpoint   string
	Bridge     common.Address
	StartBlock *big.Int
}

func ParseChainConfig(reLayerAddress string) {
	ChainCfg = map[int]Chain{}
	for i := 0; i < len(Config.Chains); i++ {
		if Config.Chains[i].StartBlock.Int64() == 0 {
			err, block := GetBlockStore(Config.Chains[i].Name, reLayerAddress)
			if err != nil {
				log.Panicln(err)
			}
			Config.Chains[i].StartBlock = block
		}

		ChainCfg[Config.Chains[i].Id] = Chain{
			Config.Chains[i].Name,
			Config.Chains[i].Id,
			Config.Chains[i].ChainId,
			Config.Chains[i].VoteChain,
			Config.Chains[i].Endpoint,
			common.HexToAddress(Config.Chains[i].Bridge),
			Config.Chains[i].StartBlock,
		}

		CreateKeyStoreIfNotExists(Config.Chains[i].Name, reLayerAddress)

		if Config.Chains[i].VoteChain {
			EngineCfg = Engine{
				Config.Engine.Name,
				Config.Engine.ChainId,
				Config.Engine.Endpoint,
				big.NewInt(Config.Chains[i].StartBlock.Int64()),
				common.HexToAddress(Config.Engine.ManagerAddress),
			}
		}

		ReLayerNum++
	}

}

func CreateKeyStoreIfNotExists(chainName, reLayerAddress string) {
	ksFile := keyStore.GetCurrentAbsPathByCaller() + "/" + chainName + "-" + reLayerAddress
	exists, err := utils.FileExists(ksFile)
	if err != nil {
		log.Panicln(err)
	}
	if !exists {
		f, err := os.Create(ksFile)
		defer f.Close()
		if err != nil {
			log.Panicln(err)
		}
	}
}

func GetBlockStore(chainName, reLayerAddress string) (error, *big.Int) {
	ksFile := keyStore.GetCurrentAbsPathByCaller() + "/" + chainName + "-" + reLayerAddress

	ks, err := os.ReadFile(ksFile)
	if err != nil {
		return err, big.NewInt(0)
	}
	return nil, big.NewInt(utils.StringToInt64(string(ks)))

}
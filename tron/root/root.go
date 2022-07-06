package root

import (
	"bridge-relayer/log"
	"bridge-relayer/tron/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	c "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"strings"
)

var (
	Addr                   address.TronAddress
	Signer                 string
	SignerAddress          address.TronAddress
	ContractName           string
	Verbose                bool
	dryRun                 bool
	NoWait                 bool
	UseLedgerWallet        bool
	NoPrettyOutput         bool
	UserProvidesPassphrase bool
	PassphraseFilePath     string
	DefaultKeystoreDir     string
	Node                   string
	KeyStoreDir            string
	GivenFilePath          string
	Timeout                uint32
	WithTLS                bool
	ApiKey                 string
	Conn                   *client.GrpcClient
)

func Opts(ctlr *transaction.Controller) {
	if dryRun {
		ctlr.Behavior.DryRun = true
	}
	if UseLedgerWallet {
		ctlr.Behavior.SigningImpl = transaction.Ledger
	}
	if NoWait {
		ctlr.Behavior.ConfirmationWaitTime = 0
	} else if Timeout > 0 {
		ctlr.Behavior.ConfirmationWaitTime = Timeout
	}
}

func init() {
	ApiKey = "609cfb3a-dfe7-4286-8324-5e44d905dcd9"
	//Node = "grpc.trongrid.io:50051"
	Node = "grpc.shasta.trongrid.io:50051"
	Signer = "TX7VatsLxHP9VhwxjtraQmbsDaBNf8ykpW"
	//Signer = "TUzWxrmFa8mvQQwG4HqB3Cv3W7rgWTamzb"
	//Signer = "41e7ebcff07619ed7956b79b9a8c13d31deb5f8ac1"
	Passphrase = "ware"
	NoWait = false
	Timeout = 20

	//c.DefaultConfigDirName = "Desktop/workspace/golang/bridge-ralayer/tron/keystore"
	c.DefaultConfigDirName = "Desktop\\bridge-ralayer\\tron\\keystore"

	//----
	//if Verbose {
	//	common.EnableAllVerbose()
	//}
	switch URLcomponents := strings.Split(Node, ":"); len(URLcomponents) {
	case 1:
		Node = Node + ":50051"
	}
	Conn = client.NewGrpcClient(Node)

	// load grpc options
	opts := make([]grpc.DialOption, 0)
	if WithTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// set API
	err := Conn.SetAPIKey(ApiKey)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	if err := Conn.Start(opts...); err != nil {
		log.Logger.Error(err.Error())
		return
	}

	if len(Signer) > 0 {
		var err error
		if SignerAddress, err = address.FindAddress(Signer); err != nil {
			log.Logger.Error(err.Error())
			return
		}
	}

	Passphrase, err = GetPassphrase()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	if len(DefaultKeystoreDir) > 0 {
		// set default directory
		store.SetDefaultLocation(DefaultKeystoreDir)
	}

}

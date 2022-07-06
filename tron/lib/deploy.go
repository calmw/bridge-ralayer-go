package lib

import "C"
import (
	"bridge-relayer/binding/bridge"
	"bridge-relayer/log"
	"bridge-relayer/tron/root"
	"encoding/json"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/contract"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"path"
)

var (
	AbiSTR     string
	AbiFile    string
	BcSTR      string
	BcFile     string
	feeLimit   int64
	curPercent int64
	oeLimit    int64
)

func initDeploy() {
	AbiFile = bridge.GetCurrentAbsPathByCaller() + "/bridge.json"
	BcFile = bridge.GetCurrentAbsPathByCaller() + "/bridge.bin"
	feeLimit = 8000000000
	curPercent = 50
	oeLimit = 8000000000
}

func Deploy(contractName string) error {
	initDeploy()
	root.ContractName = contractName
	if AbiFile != "" {
		abiBytes, err := ioutil.ReadFile(AbiFile)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("cannot read ABI file: %s %v", AbiFile, err)
		}
		AbiSTR = string(abiBytes)
	} else {
		return fmt.Errorf("no ABI string or ABI file specified")
	}
	ABI, err := contract.JSONtoABI(AbiSTR)
	if err != nil {
		return fmt.Errorf("cannot parse ABI: %v", err)
	}

	if BcFile != "" {
		bcBytes, err := ioutil.ReadFile(BcFile)
		if err != nil {
			return fmt.Errorf("cannot read Bytecode file: %s %v", BcFile, err)
		}
		BcSTR = string(bcBytes)
	} else {
		return fmt.Errorf("no Bytecode string or Bytecode file specified")
	}

	if root.SignerAddress.String() == "" {
		return fmt.Errorf("no signer specified")
	}
	if root.ContractName == "" {
		return fmt.Errorf("no contract name")
	}

	// TODO: add constructor arguments
	tx, err := root.Conn.DeployContract(root.SignerAddress.Address, root.ContractName,
		ABI, BcSTR, feeLimit, curPercent, oeLimit)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	var ctrlr *transaction.Controller
	//if root.UseLedgerWallet {
	//	account := keystore.Account{Address: root.SignerAddress.GetAddress()}
	//	ctrlr = transaction.NewController(root.Conn, nil, &account, tx.Transaction, root.Opts)
	//} else {
	uDir, _ := homedir.Dir()
	keystorePath := path.Join(
		uDir,
		common.DefaultConfigDirName,
		common.DefaultConfigAccountAliasesDirName,
	)

	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	sender, err := address.Base58ToAddress(root.SignerAddress.String())
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	account, err := ks.Find(keystore.Account{Address: sender})
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	if unlockError := ks.Unlock(account, root.Passphrase); unlockError != nil {
		log.Logger.Error(unlockError.Error())
		return unlockError
	}

	ctrlr = transaction.NewController(root.Conn, ks, &account, tx.Transaction, root.Opts)
	//}
	if err = ctrlr.ExecuteTransaction(); err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	if root.NoPrettyOutput {
		fmt.Println(tx)
		return nil
	}
	fmt.Println(ctrlr.Receipt.ContractAddress)

	addrResult := address.Address(ctrlr.Receipt.ContractAddress).String()

	result := make(map[string]interface{})
	result["txID"] = common.BytesToHexString(tx.GetTxid())
	result["blockNumber"] = ctrlr.Receipt.BlockNumber
	result["message"] = string(ctrlr.Result.Message)
	result["contractAddress"] = addrResult
	result["success"] = ctrlr.GetResultError() == nil
	result["resMessage"] = string(ctrlr.Receipt.ResMessage)
	result["receipt"] = map[string]interface{}{
		"fee":               ctrlr.Receipt.Fee,
		"energyFee":         ctrlr.Receipt.Receipt.EnergyFee,
		"energyUsage":       ctrlr.Receipt.Receipt.EnergyUsage,
		"originEnergyUsage": ctrlr.Receipt.Receipt.OriginEnergyUsage,
		"energyUsageTotal":  ctrlr.Receipt.Receipt.EnergyUsageTotal,
		"netFee":            ctrlr.Receipt.Receipt.NetFee,
		"netUsage":          ctrlr.Receipt.Receipt.NetUsage,
	}

	asJSON, _ := json.Marshal(result)
	fmt.Println(common.JSONPrettyFormat(string(asJSON)))

	return nil
}

package trigger

import (
	"bridge-relayer/log"
	"bridge-relayer/tron/root"
	"encoding/json"
	"fmt"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/mitchellh/go-homedir"
	"math/big"
	"path"
)

func TriggerConstantContract(contractAddress, method, jsonString string) error {
	if root.SignerAddress.String() == "" {
		return fmt.Errorf("no signer specified")
	}

	tx, err := root.Conn.TriggerConstantContract(
		root.SignerAddress.String(),
		contractAddress,
		method,
		jsonString,
	)
	fmt.Println(err, 33)
	if err != nil {
		return err
	}

	cResult := tx.GetConstantResult()

	fmt.Println(cResult, 11)
	if root.NoPrettyOutput {
		fmt.Println(cResult)
		return nil
	}

	result := make(map[string]interface{})
	//TODO: parse based on contract ABI
	result["Result"] = common.ToHex(cResult[0])

	asJSON, _ := json.Marshal(result)
	fmt.Println(common.JSONPrettyFormat(string(asJSON)))

	return nil
}

func TriggerContract(contractAddress, method, jsonString string, feeLimit, callValue int64, tTokenID string, tTokenAmount int64) error {
	if root.SignerAddress.String() == "" {
		return fmt.Errorf("no signer specified")
	}
	// get amount
	valueInt := int64(0)
	if callValue > 0 {
		valueInt = callValue
	}
	tokenInt := int64(0)
	//if tTokenAmount > 0 {
	//	// get token info
	//	info, err := root.Conn.GetAssetIssueByID(tTokenID)
	//	if err != nil {
	//		return err
	//	}
	//	tokenInt = int64(tAmount * math.Pow10(int(info.Precision)))
	//}

	tx, err := root.Conn.TriggerContract(
		root.SignerAddress.String(),
		contractAddress,
		method,
		jsonString,
		feeLimit,
		valueInt,
		tTokenID,
		tokenInt,
	)
	if err != nil {
		return err
	}

	var ctrlr *transaction.Controller
	//root.Passphrase="ware"
	//ks, acct, err := store.UnlockedKeystore(root.SignerAddress.String(), root.Passphrase)
	//fmt.Println(err,9,root.SignerAddress.String(),8, root.Passphrase)
	//fmt.Println(678697,root.Passphrase)
	//if err != nil {
	//	return err
	//}

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

	if err = ctrlr.ExecuteTransaction(); err != nil {
		return err
	}

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

func TriggerContract2(contractAddress, method string, packed []byte, feeLimit, callValue int64, tTokenID string, tTokenAmount int64) error {
	if root.SignerAddress.String() == "" {
		return fmt.Errorf("no signer specified")
	}
	// get amount
	valueInt := int64(0)
	if callValue > 0 {
		valueInt = callValue
	}
	tokenInt := int64(0)
	//if tTokenAmount > 0 {
	//	// get token info
	//	info, err := root.Conn.GetAssetIssueByID(tTokenID)
	//	if err != nil {
	//		return err
	//	}
	//	tokenInt = int64(tAmount * math.Pow10(int(info.Precision)))
	//}

	tx, err := root.Conn.TriggerContract2(
		root.SignerAddress.String(),
		contractAddress,
		method,
		packed,
		feeLimit,
		valueInt,
		tTokenID,
		tokenInt,
	)
	if err != nil {
		return err
	}

	var ctrlr *transaction.Controller
	//root.Passphrase="ware"
	//ks, acct, err := store.UnlockedKeystore(root.SignerAddress.String(), root.Passphrase)
	//fmt.Println(err,9,root.SignerAddress.String(),8, root.Passphrase)
	//fmt.Println(678697,root.Passphrase)
	//if err != nil {
	//	return err
	//}

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

	if err = ctrlr.ExecuteTransaction(); err != nil {
		return err
	}

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

type ExecuteRequest struct {
	ResourceID    [32]byte `json:"bytes32"`
	SourceChainId uint32   `json:"bytes"`
	SourceNonce   *big.Int `json:"uint256"`
	TargetChainId uint32   `json:"uint32"`
	//Target        address.Address `json:"address"`
	Target     ethCommon.Address `json:"address"`
	MessageId  [32]byte          `json:"bytes32"`
	DataHash   [32]byte          `json:"bytes32"`
	Data       []byte            `json:"bytes"`
	Signatures [][]byte          `json:"bytes[]"`
}

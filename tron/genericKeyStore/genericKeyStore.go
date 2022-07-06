package genericKeyStore

import (
	"bridge-relayer/log"
	"bridge-relayer/tron/root"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/mitchellh/go-homedir"
	"path"
)

// GenericKeysFromPrivateKey 通过已有账户私钥生成keystone文件
// 81290e5eac98cb77f265ff0cd92b11619ba59d95404c14dc1c66b4e42c0aca18
func GenericKeysFromPrivateKey(privateKey string) {
	uDir, _ := homedir.Dir()
	keystorePath := path.Join(
		uDir,
		common.DefaultConfigDirName,
		common.DefaultConfigAccountAliasesDirName,
	)
	fmt.Println(keystorePath)
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	ReLayerPrivateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	ecdsa, err := ks.ImportECDSA(ReLayerPrivateKeyEcdsa, root.Passphrase)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	fmt.Println(ecdsa)
}

// GenericKeys 生成keystone文件
func GenericKeys() {
	uDir, _ := homedir.Dir()
	keystorePath := path.Join(
		uDir,
		common.DefaultConfigDirName,
		common.DefaultConfigAccountAliasesDirName,
	)
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	newAccount, err := ks.NewAccount(root.Passphrase)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	fmt.Println(newAccount.Address.String())
}

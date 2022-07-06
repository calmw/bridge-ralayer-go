package lib

import (
	"bridge-relayer/tron/genericKeyStore"
)

func generic() {
	genericKeyStore.GenericKeys()
	genericKeyStore.GenericKeysFromPrivateKey("81290e5eac98cb77f265ff0cd92b11619ba59d95404c14dc1c66b4e42c0aca18")
}

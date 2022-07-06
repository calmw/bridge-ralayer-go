package root

import (
	"fmt"
	c "github.com/fbsobreira/gotron-sdk/pkg/common"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"strings"
)

const (
	seedPhraseWarning = "**Important** write this seed phrase in a safe place, " +
		"it is the only way to recover your account if you ever forget your password\n\n"
)

var (
	quietImport         bool
	recoverFromMnemonic bool
	Passphrase          string
	blsFilePath         string
	blsShardID          uint32
	blsCount            uint32
	ppPrompt            = fmt.Sprintf(
		"prompt for passphrase, otherwise use default passphrase: \"`%s`\"", c.DefaultPassphrase,
	)
)

func GetPassphrase() (string, error) {
	if PassphraseFilePath != "" {
		if _, err := os.Stat(PassphraseFilePath); os.IsNotExist(err) {
			return "", fmt.Errorf("passphrase file not found at `%s`", PassphraseFilePath)
		}
		dat, err := ioutil.ReadFile(PassphraseFilePath)
		if err != nil {
			return "", err
		}
		pw := strings.TrimSuffix(string(dat), "\n")
		return pw, nil
	} else if UserProvidesPassphrase {
		fmt.Println("Enter passphrase:")
		pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", err
		}
		return string(pass), nil
	} else {
		return c.DefaultPassphrase, nil
	}
}

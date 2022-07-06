package address

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"github.com/pkg/errors"
)

type TronAddress struct {
	Address string
}

func (t TronAddress) String() string {
	return t.Address
}

func (t *TronAddress) Set(s string) error {
	_, err := address.Base58ToAddress(s)
	if err != nil {
		return errors.Wrap(err, "not a valid one address")
	}
	t.Address = s
	return nil
}

func (t *TronAddress) GetAddress() address.Address {
	addr, err := address.Base58ToAddress(t.Address)
	if err != nil {
		return nil
	}
	return addr
}

func (t TronAddress) Type() string {
	return "tron-address"
}

func ValidateAddress(address string) error {
	// Check if input valid one address
	Addr, err := FindAddress(address)
	fmt.Println(Addr, err)
	return err
}

func FindAddress(value string) (TronAddress, error) {
	// Check if input valid one res
	res := TronAddress{}
	if err := res.Set(value); err != nil {
		// Check if input is valid account name
		if acc, err := store.AddressFromAccountName(value); err == nil {
			res.Address = acc
			return res, nil
		}
		return res, fmt.Errorf("invalid res/Invalid account name: %s", value)
	}
	return res, nil
}

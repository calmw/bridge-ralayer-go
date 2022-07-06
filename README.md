## Installation

```shell
install golang
```

## Building

`make build`

## Start bridge

``` shell
./cmd/relayer run --privateKey 8699040b13da6c1994f97bef8d2fe458bf5c23e6ca5a97d45bd4663eaf90b856 --address 0x3bd1a4c59b575eC77dDBd9c9c0a46633E5D5Bec7
```

The above parameters privateKey and address are the private key and address of the current bridge. Before executing the
command, set the address as the validator role in the deployed contract

## Deploy contract

``` shell
./cmd/relayer deploy --privateKey 8699040b13da6c1994f97bef8d2fe458bf5c23e6ca5a97d45bd4663eaf90b856 --address 0x3bd1a4c59b575eC77dDBd9c9c0a46633E5D5Bec7 --name bridge --chainId 2
```

### Contract Rpo：https://github.com/ysfinance/bridge-contracts


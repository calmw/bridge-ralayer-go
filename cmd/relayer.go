package main

import (
	"bridge-relayer/config"
	"bridge-relayer/internal/relayer"
	"bridge-relayer/services"
	"bridge-relayer/services/deploy"
	"errors"
	log "github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{}

	app.UseShortOptionHandling = true

	app.Commands = []*cli.Command{
		{
			Name:  "run",
			Usage: "run reLayer",
			Action: func(c *cli.Context) error {
				ReLayerPrivateKeyEcdsa, err := crypto.HexToECDSA(c.String("privateKey"))
				if err != nil {
					log.Error("privateKey  ", "err", err)
					return err
				}
				relayer.ThisReLayer.PrivateKey = ReLayerPrivateKeyEcdsa
				if c.String("address") == "" {
					log.Error("address  ", "err", "address empty")
					return errors.New("address empty")
				}
				relayer.ThisReLayer.Address = common.HexToAddress(c.String("address"))
				config.InitConfig(c.String("address"))
				relayer.InitReLayer()

				log.Info("Starting reLayer ...")

				go services.StartWatcher()

				engine, err := services.NewEngine()
				if err != nil {
					return err
				}
				engine.PollBlocks()

				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "privateKey",
					Usage:    "reLayer private key",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "address",
					Usage:    "reLayer address",
					Required: true,
				},
			},
		},
		{
			Name:  "deploy",
			Usage: "deploy contract",
			Action: func(c *cli.Context) error {
				ReLayerPrivateKeyEcdsa, err := crypto.HexToECDSA(c.String("privateKey"))
				if err != nil {
					log.Error("privateKey  ", "err", err)
					return err
				}
				relayer.ThisReLayer.PrivateKey = ReLayerPrivateKeyEcdsa
				if c.String("name") == "" {
					log.Error("contract name ", "err", "contract name empty")
					return errors.New("address empty")
				}
				_, has := deploy.ContractName[c.String("name")]
				if !has {
					return errors.New("this contract is not exists")
				}
				if c.String("chainId") == "" {
					log.Error("chainId ", "err", "chainId empty")
					return errors.New("chainId empty")
				}
				if c.String("address") == "" {
					log.Error("address  ", "err", "address empty")
					return errors.New("address empty")
				}
				relayer.ThisReLayer.Address = common.HexToAddress(c.String("address"))
				config.InitConfig(c.String("address"))
				relayer.InitReLayer()
				deploy.DeployContract(c.String("name"), c.String("chainId"))
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "privateKey",
					Usage:    "admin private key",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "address",
					Usage:    "reLayer address",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Usage:    "contract name",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "chainId",
					Usage:    "chain id",
					Required: true,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

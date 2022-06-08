// Copyright 2022 ReLayer Systems
package main

import (
	"bridge-relayer/config"
	"bridge-relayer/internal/relayer"
	"bridge-relayer/services"
	"errors"
	"fmt"
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

				log.Info("Starting reLayer...")

				go services.StartWatcher()

				engine, err := services.NewEngine()
				if err != nil {
					return err
				}
				engine.GetSignatureCollectedEvent()

				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "privateKey",
					Value:    "dd",
					Usage:    "reLayer private key",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "address",
					Value:    "fff",
					Usage:    "reLayer address",
					Required: true,
				},
			},
		},
		{
			Name:    "deploy",
			Aliases: []string{"d"},
			Usage:   "deploy contract",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "privateKey",
					Value:    "dd",
					Usage:    "reLayer private key",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "address",
					Value:    "fff",
					Usage:    "reLayer address",
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

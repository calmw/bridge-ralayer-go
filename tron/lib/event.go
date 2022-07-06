package lib

import (
	"bridge-relayer/log"
	"bridge-relayer/tron/event"
	"bridge-relayer/utils"
	"encoding/json"
	"errors"
)

func PollBlocks(block int64) error {
	err, res := GetEvent(block)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	for _, e := range res {
		if e.Event == "CallRequest" {

		} else if e.Event == "ConfirmedRequest" {

		}
	}

	return errors.New("no CallRequest Event")
}

func GetEvent(block int64) (error, []event.EventData) {
	blockStr := utils.Int64ToString(block)
	uri := "https://api.shasta.trongrid.io/v1/blocks/" + blockStr + "/events?only_confirmed=true&limit=200"
	res, err := utils.HttpGet(uri, map[string]string{}, nil)
	if err != nil {
		log.Logger.Error(err.Error())
		return err, nil
	}
	var eventResponse event.EventResponse
	err = json.Unmarshal(res, &eventResponse)
	if err != nil {
		log.Logger.Error(err.Error())
		return err, nil
	}
	if len(eventResponse.Data) <= 0 {
		return errors.New("no event"), nil
	}
	return nil, eventResponse.Data
}

func GetLatestBlock() (error, int) {
	uri := "https://api.shasta.trongrid.io/walletsolidity/getnowblock"
	res, err := utils.HttpGet(uri, map[string]string{}, nil)
	if err != nil {
		log.Logger.Error(err.Error())
		return err, 0
	}
	var block event.Block
	err = json.Unmarshal(res, &block)
	if err != nil {
		log.Logger.Error(err.Error())
		return err, 0
	}

	return nil, block.BlockHeader.RawData.Number
}

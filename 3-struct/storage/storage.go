package storage

import (
	"encoding/json"
	"errors"
	"main/bins"
	"os"
)

const filename = "data.json"

func SaveBins(binList bins.BinList) error {
	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func LoadBins() (bins.BinList, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return bins.BinList{}, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var binList bins.BinList
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&binList); err != nil {
		return nil, errors.New("не удалось распарсить JSON")
	}

	return binList, nil
}

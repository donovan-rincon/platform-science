package internal

import (
	"encoding/json"
	"os"
)

var input InputData

type InputData struct {
	Shipments []string `json:"shipments"`
	Drivers   []string `json:"drivers"`
}

func ReadInputFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &input)
	if err != nil {
		panic(err)
	}
}

func ReadInputFiles(shipmentsPath string, driversPath string) {
	s, err := os.ReadFile(shipmentsPath)
	if err != nil {
		panic(err)
	}
	d, err := os.ReadFile(driversPath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(s, &input)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &input)
	if err != nil {
		panic(err)
	}
}
